package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"lambda-cities-from-state/state"
	"net/http"
)

type Event struct {
	State string `json:"state"`
}

func respondWithError(err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       err.Error(),
	}
}

func respondWithSuccess(message string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       message,
	}
}

func handleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var event Event
	err := json.Unmarshal([]byte(req.Body), &event)

	if err != nil {
		return respondWithError(err), nil
	}

	selectedState := state.FromAlias(event.State)
	citiesFromState := selectedState.GetCitiesAsJson()

	return respondWithSuccess(citiesFromState), nil
}

func main() {
	lambda.Start(handleRequest)
}
