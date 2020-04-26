.EXPORT_ALL_VARIABLES:

AWS_ACCESS_KEY_ID := $(shell pass show aws/tesla-app/access-key-id)
AWS_SECRET_ACCESS_KEY := $(shell pass show aws/tesla-app/secret-access-key)

.PHONY: test
test: deploy.done
	curl --silent \
		--request POST \
		--data-raw "$(shell pass show tesla/token)" \
		--show-error \
		--dump-header - \
		"$$(terraform output url)"

.PHONY: clean
clean:
	terraform destroy
	rm init.done deploy.done tesla.zip tesla

init.done:
	terraform init
	touch $@

deploy.done: init.done main.tf tesla.zip
	@terraform apply \
		-var "tesla_access_token=$(shell pass show tesla/access-token)" \
		-var 'vehicle_id=$(shell pass show tesla/vehicle-id)' \
		-var "token=$(shell pass show tesla/token)"
	touch $@

tesla.zip: tesla
	zip $@ $<

tesla: main.go $(shell find codegen -type f)
	GOOS=linux GOARCH=amd64 go build -o $@

.PHONY: invoke
invoke:
	aws lambda invoke --region=eu-west-2 --function-name=tesla /dev/stderr

.PHONY: lambda-logs
lambda-logs:
	aws logs get-log-events \
		--log-group-name $(shell terraform output log_group_name) \
		--log-stream-name $(shell aws logs describe-log-streams --log-group-name $(shell terraform output log_group_name) --query 'logStreams[*].logStreamName' | jq '.[-1]' | sed "s/\"/'/g")

.PHONY: tesla-swagger.yml
tesla-swagger.yml:
	wget -O tesla-swagger.yml https://raw.githubusercontent.com/jonahwh/tesla-api-client/master/swagger.yml

generate-client:
	docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate \
		-i /local/tesla-swagger.yml \
		-l go \
		-o /local/codegen/swagger

.phony: get-access-token
get-access-token:
	@curl https://owner-api.teslamotors.com/oauth/token \
		--header "Content-Type: application/json" \
		--request POST \
		--data '{ \
					"grant_type": "password", \
					"client_id": "$(shell pass show tesla/client-id)", \
					"client_secret": "$(shell pass show tesla/client-secret)", \
					"email": "$(shell pass show tesla.com | grep login | cut -d' ' -f2)", \
					"password": "$(shell pass show tesla.com | head -n1)" \
				}' \
			| jq .
