package main

import (
	"context"
	"fmt"
	"log"
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
		log.Printf("failed authorization: %s", err)
		return makeResponse(http.StatusUnauthorized, err.Error()), nil
	}

	c := makeTeslaClient()
	ctx = context.WithValue(ctx, swagger.ContextAccessToken, os.Getenv("TESLA_ACCESS_TOKEN"))
	vid := os.Getenv("VEHICLE_ID")

	awake, err := wakeVehicle(c, ctx, vid)
	if err != nil {
		log.Printf("failed waking vehicle: %s", err)
		return makeResponse(http.StatusBadGateway, err.Error()), nil
	}

	if !awake {
		log.Print("tesla still sleeping...")
		return makeResponse(http.StatusUnprocessableEntity, "sleeping"), nil
	}

	ok, err := openChargeDoor(c, ctx, vid)
	if err != nil {
		log.Printf("error opening charge door: %s", err)
		return makeResponse(http.StatusBadGateway, err.Error()), nil
	}
	if !ok {
		log.Print("didn't open charge door")
		return makeResponse(http.StatusBadGateway, "didn't open charge door"), nil
	}

	return makeResponse(http.StatusOK, "successfully opened charge door"), nil
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

func openChargeDoor(c *swagger.APIClient, ctx context.Context, vid string) (bool, error) {
	ocpr, _, err := c.VehicleCommandsApi.OpenChargePort(ctx, vid)
	if err != nil {
		return false, err
	}

	return ocpr.Response.Result, nil
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
