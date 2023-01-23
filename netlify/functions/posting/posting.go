package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// Your server-side functionality
}

func main() {
	// Make the handler available for Remote Procedure Call
	lambda.Start(handler)
}
