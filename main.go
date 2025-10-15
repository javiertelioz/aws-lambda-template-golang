package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/javiertelioz/aws-lambda-golang/pkg/infrastructure/handlers"
)

func main() {
	lambda.Start(handlers.HelloHandleRequest)
}
