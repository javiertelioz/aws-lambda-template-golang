package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"

	"github.com/javiertelioz/aws-lambda-golang/pkg/application/use_cases/hello"
	"github.com/javiertelioz/aws-lambda-golang/pkg/infrastructure/sevices/logger"
)

func HelloHandleRequest(
	_ context.Context,
	request events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	loggerService := logger.NewLogger()
	loggerService.Debug(fmt.Sprintf("Request Query Params: %v", request.QueryStringParameters))

	name := request.QueryStringParameters["name"]

	if name == "" {
		name = "world"
	}

	message := hello.SayHelloUseCase(name)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       message,
	}, nil
}
