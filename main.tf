variable "tesla_access_token" {
  type = string
}

variable "token" {
  type = string
}

variable "vehicle_id" {
  type    = string
  default = ""
}

provider "aws" {
  version = "~> 2.0"
  region  = "eu-west-2"
}

resource "aws_lambda_function" "tesla" {
  function_name    = "tesla"
  filename         = "tesla.zip"
  source_code_hash = filebase64sha256("tesla.zip")
  handler          = "tesla"
  runtime          = "go1.x"
  role             = aws_iam_role.tesla.arn

  environment {
    variables = {
      TESLA_ACCESS_TOKEN = var.tesla_access_token
      TOKEN              = var.token
      VEHICLE_ID         = var.vehicle_id
    }
  }

  depends_on = [aws_iam_role_policy_attachment.lambda_logs, aws_cloudwatch_log_group.lambda]
}

resource "aws_iam_role" "tesla" {
  name = "tesla_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
  EOF
}

resource "aws_api_gateway_account" "tesla" {
  cloudwatch_role_arn = aws_iam_role.cloudwatch.arn
}

resource "aws_iam_role" "cloudwatch" {
  name = "api_gateway_cloudwatch_global"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "apigateway.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "cloudwatch" {
  name = "default"
  role = aws_iam_role.cloudwatch.id

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogGroup",
                "logs:CreateLogStream",
                "logs:DescribeLogGroups",
                "logs:DescribeLogStreams",
                "logs:PutLogEvents",
                "logs:GetLogEvents",
                "logs:FilterLogEvents"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}

resource "aws_api_gateway_rest_api" "tesla" {
  name        = "Tesla"
  description = "API to do stuff with Tesla"
}

resource "aws_api_gateway_method" "proxy_root" {
  rest_api_id   = aws_api_gateway_rest_api.tesla.id
  resource_id   = aws_api_gateway_rest_api.tesla.root_resource_id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "lambda_root" {
  rest_api_id = aws_api_gateway_rest_api.tesla.id
  resource_id = aws_api_gateway_method.proxy_root.resource_id
  http_method = aws_api_gateway_method.proxy_root.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.tesla.invoke_arn
}

resource "aws_api_gateway_deployment" "tesla" {
  rest_api_id = aws_api_gateway_rest_api.tesla.id
  stage_name  = "v1"

  depends_on = [
    aws_api_gateway_integration.lambda_root,
  ]
}

resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.tesla.function_name
  principal     = "apigateway.amazonaws.com"
}

resource "aws_cloudwatch_log_group" "lambda" {
  name              = "/aws/lambda/tesla" # avoiding dependency on aws_lambda_function.tesla.function_name
  retention_in_days = 14
}

resource "aws_iam_policy" "lambda_logging" {
  name        = "lambda_logging"
  path        = "/"
  description = "IAM policy for logging from a lambda"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*",
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.tesla.name
  policy_arn = aws_iam_policy.lambda_logging.arn
}

resource "aws_cloudwatch_log_group" "apigw" {
  name              = "API-Gateway-Execution-Logs_${aws_api_gateway_rest_api.tesla.id}/v1"
  retention_in_days = 7
}

output "url" {
  value = aws_api_gateway_deployment.tesla.invoke_url
}

output "log_group_name" {
  value = aws_cloudwatch_log_group.lambda.name
}
