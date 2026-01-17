package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/aws-lambda-golang/pkg/domain/services"
	"github.com/javiertelioz/aws-lambda-golang/pkg/infrastructure/sevices/logger"
)

type ZerologLoggerTestSuite struct {
	suite.Suite
	logger     *logger.ZerologLogger
	logOutput  *bytes.Buffer
	logMessage string
}

func TestZerologLoggerTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ZerologLoggerTestSuite))
}

func (suite *ZerologLoggerTestSuite) SetupTest() {
	suite.logOutput = &bytes.Buffer{}
	log.Logger = zerolog.New(suite.logOutput).With().Timestamp().Logger()
	suite.logger = logger.NewLogger()
	suite.logMessage = ""
}

func (suite *ZerologLoggerTestSuite) givenLogMessage(msg string) {
	suite.logMessage = msg
}

func (suite *ZerologLoggerTestSuite) whenTraceIsCalled() {
	suite.logger.Log(context.Background(), services.LevelTrace, suite.logMessage)
}

func (suite *ZerologLoggerTestSuite) whenDebugIsCalled() {
	suite.logger.Log(context.Background(), services.LevelDebug, suite.logMessage)
}

func (suite *ZerologLoggerTestSuite) whenInfoIsCalled() {
	suite.logger.Log(context.Background(), services.LevelInfo, suite.logMessage)
}

func (suite *ZerologLoggerTestSuite) whenWarnIsCalled() {
	suite.logger.Log(context.Background(), services.LevelWarn, suite.logMessage)
}

func (suite *ZerologLoggerTestSuite) whenErrorIsCalled() {
	suite.logger.Log(context.Background(), services.LevelError, suite.logMessage)
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
	suite.Contains(logEntry, "file")
	suite.Contains(logEntry, "line")
	suite.NotEmpty(logEntry["file"])
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

func (suite *ZerologLoggerTestSuite) TestLogWithRequestID_ShouldIncludeInLog() {
	// Given
	ctx := context.WithValue(context.Background(), "request_id", "req-123-456")
	suite.logMessage = "Request processed"

	// When
	suite.logger.Log(ctx, services.LevelInfo, suite.logMessage)

	// Then
	var logEntry map[string]interface{}
	err := json.Unmarshal(suite.logOutput.Bytes(), &logEntry)
	suite.NoError(err)
	suite.Equal("Request processed", logEntry["message"])
	suite.Equal("req-123-456", logEntry["request_id"])
}

func (suite *ZerologLoggerTestSuite) TestLogWithTraceID_ShouldIncludeInLog() {
	// Given
	ctx := context.WithValue(context.Background(), "trace_id", "trace-xyz-789")
	suite.logMessage = "Trace started"

	// When
	suite.logger.Log(ctx, services.LevelDebug, suite.logMessage)

	// Then
	var logEntry map[string]interface{}
	err := json.Unmarshal(suite.logOutput.Bytes(), &logEntry)
	suite.NoError(err)
	suite.Equal("Trace started", logEntry["message"])
	suite.Equal("trace-xyz-789", logEntry["trace_id"])
}

func (suite *ZerologLoggerTestSuite) TestLogWithCorrelationID_ShouldIncludeInLog() {
	// Given
	ctx := context.WithValue(context.Background(), "correlation_id", "corr-abc-def")
	suite.logMessage = "Correlated request"

	// When
	suite.logger.Log(ctx, services.LevelInfo, suite.logMessage)

	// Then
	var logEntry map[string]interface{}
	err := json.Unmarshal(suite.logOutput.Bytes(), &logEntry)
	suite.NoError(err)
	suite.Equal("Correlated request", logEntry["message"])
	suite.Equal("corr-abc-def", logEntry["correlation_id"])
}

func (suite *ZerologLoggerTestSuite) TestLogWithUserID_ShouldIncludeInLog() {
	// Given
	ctx := context.WithValue(context.Background(), "user_id", 12345)
	suite.logMessage = "User action"

	// When
	suite.logger.Log(ctx, services.LevelInfo, suite.logMessage)

	// Then
	var logEntry map[string]interface{}
	err := json.Unmarshal(suite.logOutput.Bytes(), &logEntry)
	suite.NoError(err)
	suite.Equal("User action", logEntry["message"])
	suite.Equal(float64(12345), logEntry["user_id"]) // JSON unmarshals numbers as float64
}

func (suite *ZerologLoggerTestSuite) TestLogWithAllContextValues_ShouldIncludeAll() {
	// Given
	ctx := context.Background()
	ctx = context.WithValue(ctx, "request_id", "req-001")
	ctx = context.WithValue(ctx, "trace_id", "trace-002")
	ctx = context.WithValue(ctx, "correlation_id", "corr-003")
	ctx = context.WithValue(ctx, "user_id", "user-789")
	suite.logMessage = "Complete context"

	// When
	suite.logger.Log(ctx, services.LevelInfo, suite.logMessage)

	// Then
	var logEntry map[string]interface{}
	err := json.Unmarshal(suite.logOutput.Bytes(), &logEntry)
	suite.NoError(err)
	suite.Equal("Complete context", logEntry["message"])
	suite.Equal("req-001", logEntry["request_id"])
	suite.Equal("trace-002", logEntry["trace_id"])
	suite.Equal("corr-003", logEntry["correlation_id"])
	suite.Equal("user-789", logEntry["user_id"])
}

func (suite *ZerologLoggerTestSuite) TestLogWithNilContext_ShouldNotPanic() {
	// Given
	suite.logMessage = "Nil context test"

	// When
	suite.logger.Log(nil, services.LevelInfo, suite.logMessage)

	// Then
	var logEntry map[string]interface{}
	err := json.Unmarshal(suite.logOutput.Bytes(), &logEntry)
	suite.NoError(err)
	suite.Equal("Nil context test", logEntry["message"])

	suite.NotContains(logEntry, "request_id")
	suite.NotContains(logEntry, "trace_id")
}

func (suite *ZerologLoggerTestSuite) TestLogWithoutContextValues_ShouldNotIncludeThem() {
	// Given
	ctx := context.Background()
	suite.logMessage = "Empty context"

	// When
	suite.logger.Log(ctx, services.LevelInfo, suite.logMessage)

	// Then
	var logEntry map[string]interface{}
	err := json.Unmarshal(suite.logOutput.Bytes(), &logEntry)
	suite.NoError(err)
	suite.Equal("Empty context", logEntry["message"])
	suite.NotContains(logEntry, "request_id")
	suite.NotContains(logEntry, "trace_id")
	suite.NotContains(logEntry, "correlation_id")
	suite.NotContains(logEntry, "user_id")
}

func (suite *ZerologLoggerTestSuite) TestLogWithInvalidLevel_ShouldDefaultToInfo() {
	// Given
	ctx := context.Background()
	suite.logMessage = "Invalid level test"
	invalidLevel := services.Level(999)

	// When
	suite.logger.Log(ctx, invalidLevel, suite.logMessage)

	// Then
	var logEntry map[string]interface{}
	err := json.Unmarshal(suite.logOutput.Bytes(), &logEntry)
	suite.NoError(err)
	suite.Equal("Invalid level test", logEntry["message"])
	suite.Equal("info", logEntry["level"])
	suite.Equal("info", logEntry["log_level"])
}

func (suite *ZerologLoggerTestSuite) TestLogShouldIncludeFileAndLine() {
	// Given
	ctx := context.Background()
	suite.logMessage = "Testing file and line"

	// When
	suite.logger.Log(ctx, services.LevelInfo, suite.logMessage)

	// Then
	var logEntry map[string]interface{}
	err := json.Unmarshal(suite.logOutput.Bytes(), &logEntry)
	suite.NoError(err)

	suite.Contains(logEntry, "file")
	suite.Contains(logEntry, "line")

	suite.NotEqual("", logEntry["file"])

	_, ok := logEntry["line"].(float64)
	suite.True(ok)
}
