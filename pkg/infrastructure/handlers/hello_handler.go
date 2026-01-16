package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/aws-lambda-go/events"

	"github.com/javiertelioz/aws-lambda-golang/pkg/application/use_cases/hello"
	"github.com/javiertelioz/aws-lambda-golang/pkg/infrastructure/sevices/logger"
)

const (
	// maxNameLength defines the maximum allowed length for the name parameter
	maxNameLength = 100
)

// validNamePattern matches alphanumeric characters, spaces, hyphens, apostrophes, and common international characters
var validNamePattern = regexp.MustCompile(`^[a-zA-Z0-9\s\-'áéíóúÁÉÍÓÚñÑüÜ]+$`)

// HelloHandleRequest processes AWS API Gateway requests for the hello endpoint.
// It accepts an optional 'name' query parameter and returns a personalized greeting.
//
// Query Parameters:
//   - name (optional): The name to include in the greeting. Defaults to "world" if not provided.
//     Must be alphanumeric with max length of 100 characters.
//
// Returns:
//   - APIGatewayProxyResponse with status 200 and greeting message in the body
//   - APIGatewayProxyResponse with status 400 if validation fails
//   - error is always nil in current implementation
//
// Example requests:
//
//	GET /hello?name=John  -> Response: "Hello John!"
//	GET /hello            -> Response: "Hello world!"
//	GET /hello?name=<script> -> Response: 400 Bad Request
//
// The function uses the provided context for logging and potential timeout handling.
func HelloHandleRequest(
	ctx context.Context,
	request events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	loggerService := logger.NewLogger()
	loggerService.Debug(ctx, fmt.Sprintf("Request Query Params: %v", request.QueryStringParameters))

	name := request.QueryStringParameters["name"]

	// Validate and sanitize the name parameter
	if name != "" {
		// Trim whitespace
		name = strings.TrimSpace(name)

		// Check if empty after trimming
		if name == "" {
			name = "world"
		} else {
			// Validate length
			if len(name) > maxNameLength {
				loggerService.Warn(ctx, fmt.Sprintf("Name parameter too long: %d characters", len(name)))
				return errorResponse(400, fmt.Sprintf("Name parameter too long. Maximum %d characters allowed.", maxNameLength))
			}

			// Validate characters (prevent injection attacks)
			if !validNamePattern.MatchString(name) {
				loggerService.Warn(ctx, fmt.Sprintf("Invalid characters in name parameter: %s", name))
				return errorResponse(400, "Name parameter contains invalid characters. Only letters, numbers, spaces, and hyphens are allowed.")
			}
		}
	} else {
		name = "world"
	}

	message := hello.SayHelloUseCase(name)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       message,
	}, nil
}

// errorResponse creates a standardized error response
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
