package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/english/tesla/codegen/swagger"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, evt events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if err := authorize(evt); err != nil {
		return makeResponse(http.StatusUnauthorized, err.Error()), nil
	}

	c := makeTeslaClient()
	ctx = context.WithValue(ctx, swagger.ContextAccessToken, os.Getenv("TESLA_ACCESS_TOKEN"))
	vid := os.Getenv("VEHICLE_ID")

	awake, err := wakeVehicle(c, ctx, vid)
	if err != nil {
		return makeResponse(http.StatusBadGateway, err.Error()), nil
	}

	if !awake {
		return makeResponse(http.StatusUnprocessableEntity, "sleeping"), nil
	}

	msg, err := toggleChargeDoor(c, ctx, vid)
	if err != nil {
		return makeResponse(http.StatusBadGateway, err.Error()), nil
	}

	return makeResponse(http.StatusOK, string(msg)), nil
}

func authorize(evt events.APIGatewayProxyRequest) error {
	if evt.Body != os.Getenv("TOKEN") {
		return fmt.Errorf("invalid token: %s", evt.Body)
	}

	return nil
}

func wakeVehicle(c *swagger.APIClient, ctx context.Context, vid string) (bool, error) {
	resp, _, err := c.VehicleCommandsApi.WakeUpVehicle(ctx, vid)
	if err != nil {
		return false, fmt.Errorf("error waking up vehicle: %w", err)
	}

	return resp.Response.State == "online", nil
}

func toggleChargeDoor(c *swagger.APIClient, ctx context.Context, vid string) (string, error) {
	ok, err := openChargeDoor(c, ctx, vid)
	if err != nil {
		return "", fmt.Errorf("error opening charge port: %w", err)
	}
	if ok {
		return "was closed, now open", nil
	}

	ok, reason, err := closeChargeDoor(c, ctx, vid)
	if err != nil {
		return "", fmt.Errorf("error closing charge port: %w", err)
	}
	if !ok {
		return "", fmt.Errorf("failed closing charge port: %s", reason)
	}
	return "was open, now closed", nil
}

func openChargeDoor(c *swagger.APIClient, ctx context.Context, vid string) (bool, error) {
	ocpr, _, err := c.VehicleCommandsApi.OpenChargePort(ctx, vid)
	if err != nil {
		return false, err
	}

	return ocpr.Response.Result, nil
}

func closeChargeDoor(c *swagger.APIClient, ctx context.Context, vid string) (bool, string, error) {
	resp, _, err := c.VehicleCommandsApi.CloseChargePort(ctx, vid)
	if err != nil {
		return false, "", err
	}

	return resp.Response.Result, resp.Response.Reason, nil
}

func makeTeslaClient() *swagger.APIClient {
	cfg := swagger.NewConfiguration()
	cfg.HTTPClient = &http.Client{
		Timeout: 3 * time.Second,
	}
	return swagger.NewAPIClient(cfg)
}

func makeResponse(status int, body string) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		Body: body,
		Headers: map[string]string{
			"ContentType": "text/plain",
		},
		StatusCode: status,
	}
}
