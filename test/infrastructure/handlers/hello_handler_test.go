package handlers

import (
	"context"
	"encoding/json"
	"strings"
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
	t.Parallel()
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

func (suite *HelloHandlerTestSuite) givenRequestWithNameContainingOnlySpaces() {
	suite.request.QueryStringParameters = map[string]string{"name": "   "}
}

func (suite *HelloHandlerTestSuite) givenRequestWithLongName(length int) {
	longName := strings.Repeat("a", length)
	suite.request.QueryStringParameters = map[string]string{"name": longName}
}

func (suite *HelloHandlerTestSuite) givenRequestWithInvalidName(invalidName string) {
	suite.request.QueryStringParameters = map[string]string{"name": invalidName}
}

func (suite *HelloHandlerTestSuite) whenHelloHandleRequestIsCalled() {
	suite.response, suite.err = handlers.HelloHandleRequest(suite.ctx, suite.request)
}

func (suite *HelloHandlerTestSuite) thenResponseShouldBeSuccessful() {
	suite.NoError(suite.err)
	suite.Equal(200, suite.response.StatusCode)
}

func (suite *HelloHandlerTestSuite) thenResponseShouldBeBadRequest() {
	suite.NoError(suite.err)
	suite.Equal(400, suite.response.StatusCode)
}

func (suite *HelloHandlerTestSuite) thenResponseBodyShouldBe(expectedBody string) {
	suite.Equal(expectedBody, suite.response.Body)
}

func (suite *HelloHandlerTestSuite) thenResponseBodyShouldContain(expectedText string) {
	suite.Contains(suite.response.Body, expectedText)
}

func (suite *HelloHandlerTestSuite) thenResponseShouldBeValidJSON() {
	var jsonResponse map[string]string
	err := json.Unmarshal([]byte(suite.response.Body), &jsonResponse)
	suite.NoError(err)
	suite.Contains(jsonResponse, "error")
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

func (suite *HelloHandlerTestSuite) TestValidName_WithSpaces() {
	// Given
	suite.givenRequestWithName("John Doe")

	// When
	suite.whenHelloHandleRequestIsCalled()

	// Then
	suite.thenResponseShouldBeSuccessful()
	suite.thenResponseBodyShouldBe("Hello John Doe!")
}

func (suite *HelloHandlerTestSuite) TestValidName_WithInternationalCharacters() {
	// Given
	suite.givenRequestWithName("José María")

	// When
	suite.whenHelloHandleRequestIsCalled()

	// Then
	suite.thenResponseShouldBeSuccessful()
	suite.thenResponseBodyShouldBe("Hello José María!")
}

func (suite *HelloHandlerTestSuite) TestValidName_WithHyphen() {
	// Given
	suite.givenRequestWithName("Mary-Jane")

	// When
	suite.whenHelloHandleRequestIsCalled()

	// Then
	suite.thenResponseShouldBeSuccessful()
	suite.thenResponseBodyShouldBe("Hello Mary-Jane!")
}

func (suite *HelloHandlerTestSuite) TestValidName_WithApostrophe() {
	// Given
	suite.givenRequestWithName("O'Brien")

	// When
	suite.whenHelloHandleRequestIsCalled()

	// Then
	suite.thenResponseShouldBeSuccessful()
	suite.thenResponseBodyShouldBe("Hello O'Brien!")
}

func (suite *HelloHandlerTestSuite) TestNameWithOnlySpaces_ShouldDefaultToWorld() {
	// Given
	suite.givenRequestWithNameContainingOnlySpaces()

	// When
	suite.whenHelloHandleRequestIsCalled()

	// Then
	suite.thenResponseShouldBeSuccessful()
	suite.thenResponseBodyShouldBe("Hello world!")
}

func (suite *HelloHandlerTestSuite) TestNameTooLong_ShouldReject() {
	// Given
	suite.givenRequestWithLongName(101)

	// When
	suite.whenHelloHandleRequestIsCalled()

	// Then
	suite.thenResponseShouldBeBadRequest()
	suite.thenResponseBodyShouldContain("exceeds maximum length")
	suite.thenResponseBodyShouldContain("100")
	suite.thenResponseShouldBeValidJSON()
}

func (suite *HelloHandlerTestSuite) TestInvalidCharacters_ScriptTag_ShouldReject() {
	// Given
	suite.givenRequestWithInvalidName("<script>alert('xss')</script>")

	// When
	suite.whenHelloHandleRequestIsCalled()

	// Then
	suite.thenResponseShouldBeBadRequest()
	suite.thenResponseBodyShouldContain("contains invalid characters")
	suite.thenResponseShouldBeValidJSON()
}

func (suite *HelloHandlerTestSuite) TestInvalidCharacters_SQLInjection_ShouldReject() {
	// Given
	suite.givenRequestWithInvalidName("'; DROP TABLE users--")

	// When
	suite.whenHelloHandleRequestIsCalled()

	// Then
	suite.thenResponseShouldBeBadRequest()
	suite.thenResponseBodyShouldContain("contains invalid characters")
	suite.thenResponseShouldBeValidJSON()
}

func (suite *HelloHandlerTestSuite) TestInvalidCharacters_SpecialSymbols_ShouldReject() {
	// Given
	suite.givenRequestWithInvalidName("John@Doe")

	// When
	suite.whenHelloHandleRequestIsCalled()

	// Then
	suite.thenResponseShouldBeBadRequest()
	suite.thenResponseBodyShouldContain("contains invalid characters")
	suite.thenResponseShouldBeValidJSON()
}

func (suite *HelloHandlerTestSuite) TestInvalidCharacters_PathTraversal_ShouldReject() {
	// Given
	suite.givenRequestWithInvalidName("../../../etc/passwd")

	// When
	suite.whenHelloHandleRequestIsCalled()

	// Then
	suite.thenResponseShouldBeBadRequest()
	suite.thenResponseBodyShouldContain("contains invalid characters")
	suite.thenResponseShouldBeValidJSON()
}
