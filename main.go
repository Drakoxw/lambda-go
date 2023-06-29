package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type myEvent struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

func handleHello(ctx context.Context, ev events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body myEvent

	res := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET,HEAD,OPTIONS,POST",
		},
	}
	err := json.Unmarshal([]byte(ev.Body), &body)
	if err != nil {
		res.StatusCode = http.StatusBadRequest
		return res, err
	}

	res.Body = fmt.Sprintf(`{"message": "Hola %s %s"}`, body.Name, body.Lastname)
	res.StatusCode = http.StatusOK

	return res, nil
}

func main() {
	lambda.Start(handleHello)
}
