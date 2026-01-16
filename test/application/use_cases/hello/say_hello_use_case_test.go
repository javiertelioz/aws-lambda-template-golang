package hello

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/aws-lambda-golang/pkg/application/use_cases/hello"
)

type SayHelloUseCaseTestSuite struct {
	suite.Suite
	name   string
	result string
}

func TestSayHelloUseCaseTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(SayHelloUseCaseTestSuite))
}

func (suite *SayHelloUseCaseTestSuite) SetupTest() {
	suite.result = "Hello Joe!"
}

func (suite *SayHelloUseCaseTestSuite) givenName(name string) {
	suite.name = name
}

func (suite *SayHelloUseCaseTestSuite) whenSayHelloUseCaseIsCalled() {
	suite.result = hello.SayHelloUseCase(suite.name)
}

func (suite *SayHelloUseCaseTestSuite) thenSayHello() {
	suite.Equal(suite.result, fmt.Sprintf("Hello %s!", suite.name))
}

func (suite *SayHelloUseCaseTestSuite) TestSayHelloUseCase() {
	// Given
	suite.givenName("Joe")

	// When
	suite.whenSayHelloUseCaseIsCalled()

	// Then
	suite.thenSayHello()
}
