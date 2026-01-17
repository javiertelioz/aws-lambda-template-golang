package hello

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/aws-lambda-golang/pkg/application/use_cases/hello"
)

type SayHelloUseCaseTestSuite struct {
	suite.Suite
	name   string
	result string
	err    error
}

func TestSayHelloUseCaseTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(SayHelloUseCaseTestSuite))
}

func (suite *SayHelloUseCaseTestSuite) SetupTest() {
	suite.name = ""
	suite.result = ""
	suite.err = nil
}

func (suite *SayHelloUseCaseTestSuite) givenValidName(name string) {
	suite.name = name
}

func (suite *SayHelloUseCaseTestSuite) givenEmptyName() {
	suite.name = ""
}

func (suite *SayHelloUseCaseTestSuite) givenNameWithSpaces() {
	suite.name = "   "
}

func (suite *SayHelloUseCaseTestSuite) givenLongName(length int) {
	suite.name = strings.Repeat("a", length)
}

func (suite *SayHelloUseCaseTestSuite) givenInvalidName(name string) {
	suite.name = name
}

func (suite *SayHelloUseCaseTestSuite) whenSayHelloUseCaseIsCalled() {
	suite.result, suite.err = hello.SayHelloUseCase(suite.name)
}

func (suite *SayHelloUseCaseTestSuite) thenShouldReturnGreeting(expected string) {
	suite.NoError(suite.err)
	suite.Equal(expected, suite.result)
}

func (suite *SayHelloUseCaseTestSuite) thenShouldReturnError(expectedErr error) {
	suite.Error(suite.err)
	suite.True(errors.Is(suite.err, expectedErr))
	suite.Empty(suite.result)
}

func (suite *SayHelloUseCaseTestSuite) TestSayHelloWithValidName() {
	// Given
	suite.givenValidName("Joe")

	// When
	suite.whenSayHelloUseCaseIsCalled()

	// Then
	suite.thenShouldReturnGreeting("Hello Joe!")
}

func (suite *SayHelloUseCaseTestSuite) TestSayHelloWithNameContainingSpaces() {
	// Given
	suite.givenValidName("John Doe")

	// When
	suite.whenSayHelloUseCaseIsCalled()

	// Then
	suite.thenShouldReturnGreeting("Hello John Doe!")
}

func (suite *SayHelloUseCaseTestSuite) TestSayHelloWithInternationalCharacters() {
	// Given
	suite.givenValidName("José María")

	// When
	suite.whenSayHelloUseCaseIsCalled()

	// Then
	suite.thenShouldReturnGreeting("Hello José María!")
}

func (suite *SayHelloUseCaseTestSuite) TestSayHelloWithHyphen() {
	// Given
	suite.givenValidName("Mary-Jane")

	// When
	suite.whenSayHelloUseCaseIsCalled()

	// Then
	suite.thenShouldReturnGreeting("Hello Mary-Jane!")
}

func (suite *SayHelloUseCaseTestSuite) TestSayHelloWithApostrophe() {
	// Given
	suite.givenValidName("O'Brien")

	// When
	suite.whenSayHelloUseCaseIsCalled()

	// Then
	suite.thenShouldReturnGreeting("Hello O'Brien!")
}

func (suite *SayHelloUseCaseTestSuite) TestSayHelloWithEmptyName_ShouldDefaultToWorld() {
	// Given
	suite.givenEmptyName()

	// When
	suite.whenSayHelloUseCaseIsCalled()

	// Then
	suite.thenShouldReturnGreeting("Hello world!")
}

func (suite *SayHelloUseCaseTestSuite) TestSayHelloWithOnlySpaces_ShouldDefaultToWorld() {
	// Given
	suite.givenNameWithSpaces()

	// When
	suite.whenSayHelloUseCaseIsCalled()

	// Then
	suite.thenShouldReturnGreeting("Hello world!")
}

func (suite *SayHelloUseCaseTestSuite) TestSayHelloWithNameTooLong_ShouldReturnError() {
	// Given
	suite.givenLongName(101)

	// When
	suite.whenSayHelloUseCaseIsCalled()

	// Then
	suite.thenShouldReturnError(hello.ErrNameTooLong)
}

func (suite *SayHelloUseCaseTestSuite) TestSayHelloWithInvalidCharacters_ScriptTag_ShouldReturnError() {
	// Given
	suite.givenInvalidName("<script>alert('xss')</script>")

	// When
	suite.whenSayHelloUseCaseIsCalled()

	// Then
	suite.thenShouldReturnError(hello.ErrInvalidCharacters)
}

func (suite *SayHelloUseCaseTestSuite) TestSayHelloWithInvalidCharacters_SQLInjection_ShouldReturnError() {
	// Given
	suite.givenInvalidName("'; DROP TABLE users--")

	// When
	suite.whenSayHelloUseCaseIsCalled()

	// Then
	suite.thenShouldReturnError(hello.ErrInvalidCharacters)
}

func (suite *SayHelloUseCaseTestSuite) TestSayHelloWithInvalidCharacters_SpecialSymbols_ShouldReturnError() {
	// Given
	suite.givenInvalidName("John@Doe")

	// When
	suite.whenSayHelloUseCaseIsCalled()

	// Then
	suite.thenShouldReturnError(hello.ErrInvalidCharacters)
}

func (suite *SayHelloUseCaseTestSuite) TestSayHelloWithInvalidCharacters_PathTraversal_ShouldReturnError() {
	// Given
	suite.givenInvalidName("../../../etc/passwd")

	// When
	suite.whenSayHelloUseCaseIsCalled()

	// Then
	suite.thenShouldReturnError(hello.ErrInvalidCharacters)
}
