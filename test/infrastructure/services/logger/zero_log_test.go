package logger

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/aws-lambda-golang/pkg/infrastructure/sevices/logger"
)

type ZerologLoggerTestSuite struct {
	suite.Suite
	logger     *logger.ZerologLogger
	logOutput  *bytes.Buffer
	logMessage string
}

func TestZerologLoggerTestSuite(t *testing.T) {
	suite.Run(t, new(ZerologLoggerTestSuite))
}

func (suite *ZerologLoggerTestSuite) SetupTest() {
	suite.logOutput = &bytes.Buffer{}
	log.Logger = zerolog.New(suite.logOutput).With().Timestamp().Logger()
	suite.logger = logger.NewLogger().(*logger.ZerologLogger)
	suite.logMessage = ""
}

func (suite *ZerologLoggerTestSuite) givenLogMessage(msg string) {
	suite.logMessage = msg
}

func (suite *ZerologLoggerTestSuite) whenTraceIsCalled() {
	suite.logger.Trace(suite.logMessage)
}

func (suite *ZerologLoggerTestSuite) whenDebugIsCalled() {
	suite.logger.Debug(suite.logMessage)
}

func (suite *ZerologLoggerTestSuite) whenInfoIsCalled() {
	suite.logger.Info(suite.logMessage)
}

func (suite *ZerologLoggerTestSuite) whenWarnIsCalled() {
	suite.logger.Warn(suite.logMessage)
}

func (suite *ZerologLoggerTestSuite) whenErrorIsCalled() {
	suite.logger.Error(suite.logMessage)
}

func (suite *ZerologLoggerTestSuite) thenLogShouldContainMessage() {
	var logEntry map[string]interface{}
	err := json.Unmarshal(suite.logOutput.Bytes(), &logEntry)
	suite.NoError(err)
	suite.Equal(suite.logMessage, logEntry["message"])
}

func (suite *ZerologLoggerTestSuite) thenLogShouldContainLevel(level string) {
	var logEntry map[string]interface{}
	err := json.Unmarshal(suite.logOutput.Bytes(), &logEntry)
	suite.NoError(err)
	suite.Equal(level, logEntry["level"])
}

func (suite *ZerologLoggerTestSuite) thenLogShouldContainLocation() {
	var logEntry map[string]interface{}
	err := json.Unmarshal(suite.logOutput.Bytes(), &logEntry)
	suite.NoError(err)
	suite.Contains(logEntry, "loc")
	suite.NotEmpty(logEntry["loc"])
}

func (suite *ZerologLoggerTestSuite) TestTraceLog() {
	// Given
	suite.givenLogMessage("This is a trace message")

	// When
	suite.whenTraceIsCalled()

	// Then
	suite.thenLogShouldContainMessage()
	suite.thenLogShouldContainLevel("trace")
	suite.thenLogShouldContainLocation()
}

func (suite *ZerologLoggerTestSuite) TestDebugLog() {
	// Given
	suite.givenLogMessage("This is a debug message")

	// When
	suite.whenDebugIsCalled()

	// Then
	suite.thenLogShouldContainMessage()
	suite.thenLogShouldContainLevel("debug")
	suite.thenLogShouldContainLocation()
}

func (suite *ZerologLoggerTestSuite) TestInfoLog() {
	// Given
	suite.givenLogMessage("This is an info message")

	// When
	suite.whenInfoIsCalled()

	// Then
	suite.thenLogShouldContainMessage()
	suite.thenLogShouldContainLevel("info")
	suite.thenLogShouldContainLocation()
}

func (suite *ZerologLoggerTestSuite) TestWarnLog() {
	// Given
	suite.givenLogMessage("This is a warning message")

	// When
	suite.whenWarnIsCalled()

	// Then
	suite.thenLogShouldContainMessage()
	suite.thenLogShouldContainLevel("warn")
	suite.thenLogShouldContainLocation()
}

func (suite *ZerologLoggerTestSuite) TestErrorLog() {
	// Given
	suite.givenLogMessage("This is an error message")

	// When
	suite.whenErrorIsCalled()

	// Then
	suite.thenLogShouldContainMessage()
	suite.thenLogShouldContainLevel("error")
	suite.thenLogShouldContainLocation()
}
