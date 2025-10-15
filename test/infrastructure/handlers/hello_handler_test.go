package handlers

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/aws-lambda-golang/pkg/infrastructure/handlers"
)

type HelloHandlerTestSuite struct {
	suite.Suite
	ctx      context.Context
	request  events.APIGatewayProxyRequest
	response events.APIGatewayProxyResponse
	err      error
}

func TestHelloHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HelloHandlerTestSuite))
}

func (suite *HelloHandlerTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.request = events.APIGatewayProxyRequest{}
	suite.err = nil
}

func (suite *HelloHandlerTestSuite) givenRequestWithName(name string) {
	suite.request.QueryStringParameters = map[string]string{
		"name": name,
	}
}

func (suite *HelloHandlerTestSuite) givenRequestWithoutName() {
	suite.request.QueryStringParameters = map[string]string{}
}

func (suite *HelloHandlerTestSuite) whenHelloHandleRequestIsCalled() {
	suite.response, suite.err = handlers.HelloHandleRequest(suite.ctx, suite.request)
}

func (suite *HelloHandlerTestSuite) thenResponseShouldBeSuccessful() {
	suite.NoError(suite.err)
	suite.Equal(200, suite.response.StatusCode)
}

func (suite *HelloHandlerTestSuite) thenResponseBodyShouldBe(expectedBody string) {
	suite.Equal(expectedBody, suite.response.Body)
}

func (suite *HelloHandlerTestSuite) TestHelloHandlerWithName() {
	// Given
	suite.givenRequestWithName("Joe")

	// When
	suite.whenHelloHandleRequestIsCalled()

	// Then
	suite.thenResponseShouldBeSuccessful()
	suite.thenResponseBodyShouldBe("Hello Joe!")
}

func (suite *HelloHandlerTestSuite) TestHelloHandlerWithoutName() {
	// Given
	suite.givenRequestWithoutName()

	// When
	suite.whenHelloHandleRequestIsCalled()

	// Then
	suite.thenResponseShouldBeSuccessful()
	suite.thenResponseBodyShouldBe("Hello world!")
}
