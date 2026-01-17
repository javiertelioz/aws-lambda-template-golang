package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"

	"github.com/javiertelioz/aws-lambda-golang/pkg/application/use_cases/hello"
	"github.com/javiertelioz/aws-lambda-golang/pkg/domain/services"
	"github.com/javiertelioz/aws-lambda-golang/pkg/infrastructure/sevices/logger"
)

// HelloHandleRequest processes AWS API Gateway requests for the hello endpoint.
// It delegates input validation and business logic to the use case layer.
//
// Query Parameters:
//   - name (optional): The name to include in the greeting.
//
// Returns:
//   - APIGatewayProxyResponse with status 200 and greeting message in the body
//   - APIGatewayProxyResponse with status 400 if validation fails
//
// Example requests:
//
//	GET /hello?name=John     -> 200: "Hello John!"
//	GET /hello               -> 200: "Hello world!"
//	GET /hello?name=<script> -> 400: Validation error
func HelloHandleRequest(
	ctx context.Context,
	request events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	if request.RequestContext.RequestID != "" {
		ctx = context.WithValue(ctx, "request_id", request.RequestContext.RequestID)
	}

	loggerService := logger.NewLogger()
	loggerService.Log(ctx, services.LevelDebug, "Request received",
		services.Field{Key: "query_params", Value: request.QueryStringParameters},
		services.Field{Key: "http_method", Value: request.HTTPMethod},
		services.Field{Key: "path", Value: request.Path},
	)

	name := request.QueryStringParameters["name"]
	message, err := hello.SayHelloUseCase(name)
	if err != nil {
		loggerService.Log(ctx, services.LevelWarn, "Validation failed",
			services.Field{Key: "name", Value: name},
			services.Field{Key: "error", Value: err.Error()},
		)

		return mapErrorToResponse(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       message,
	}, nil
}

func mapErrorToResponse(err error) (events.APIGatewayProxyResponse, error) {
	var statusCode int
	var message string

	switch {
	case errors.Is(err, hello.ErrNameTooLong):
		statusCode = 400
		message = fmt.Sprintf("%s. Maximum %d characters allowed.", err.Error(), hello.MaxNameLength)
	case errors.Is(err, hello.ErrInvalidCharacters):
		statusCode = 400
		message = err.Error() + ". Only letters, numbers, spaces, hyphens, and apostrophes are allowed."
	default:
		statusCode = 500
		message = "Internal server error"
	}

	return errorResponse(statusCode, message)
}

func errorResponse(statusCode int, message string) (events.APIGatewayProxyResponse, error) {
	errorBody := map[string]string{
		"error":  message,
		"status": fmt.Sprintf("%d", statusCode),
	}

	body, _ := json.Marshal(errorBody)

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}
