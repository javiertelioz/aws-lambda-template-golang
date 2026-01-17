package logger

import (
	"context"
	"runtime"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/javiertelioz/aws-lambda-golang/pkg/domain/services"
)

type ZerologLogger struct{}

// NewLogger creates and configures a new ZerologLogger instance.
// It sets up zerolog with Unix timestamp format and stack trace support.
//
// Returns:
//   - *ZerologLogger concrete type (not interface)
//
// Following Go best practice: "Accept interfaces, return structs"
// This allows callers to use the concrete type directly or assign to an interface as needed.
//
// The logger outputs structured JSON logs to stdout with the following features:
//   - Unix timestamp format
//   - Caller location (file:line)
//   - Stack traces for error logs
//   - Structured fields support
//   - Context-based trace IDs and request IDs
func NewLogger() *ZerologLogger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return &ZerologLogger{}
}

// Log writes a log message at the specified level with optional structured fields.
// It automatically captures:
//   - Caller's file and line number for debugging
//   - Context values for distributed tracing (trace_id, request_id, correlation_id, user_id)
//   - Custom structured fields
//
// The context is used to extract metadata for log correlation across microservices
// and distributed systems. Common context keys supported:
//   - "request_id": AWS Lambda request ID or HTTP request ID
//   - "trace_id": AWS X-Ray trace ID or OpenTelemetry trace ID
//   - "correlation_id": Correlation ID for request tracking across services
//   - "user_id": Authenticated user identifier
//
// Example:
//
//	ctx := context.WithValue(ctx, "request_id", "abc-123")
//	logger.Log(ctx, services.LevelInfo, "User action",
//	    services.Field{Key: "action", Value: "login"},
//	    services.Field{Key: "ip", Value: "192.168.1.1"},
//	)
//
// This reduces code duplication by consolidating all log levels into a single implementation.
func (z *ZerologLogger) Log(ctx context.Context, level services.Level, msg string, fields ...services.Field) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	event := z.getEventForLevel(level)
	event = event.Str("file", file).Int("line", line)
	event = event.Str("log_level", level.String())

	if ctx != nil {
		if requestID := ctx.Value("request_id"); requestID != nil {
			event = event.Interface("request_id", requestID)
		}

		if traceID := ctx.Value("trace_id"); traceID != nil {
			event = event.Interface("trace_id", traceID)
		}

		if correlationID := ctx.Value("correlation_id"); correlationID != nil {
			event = event.Interface("correlation_id", correlationID)
		}

		if userID := ctx.Value("user_id"); userID != nil {
			event = event.Interface("user_id", userID)
		}
	}

	for _, field := range fields {
		event = event.Interface(field.Key, field.Value)
	}

	event.Msg(msg)
}

func (z *ZerologLogger) getEventForLevel(level services.Level) *zerolog.Event {
	switch level {
	case services.LevelTrace:
		return log.Trace()
	case services.LevelDebug:
		return log.Debug()
	case services.LevelInfo:
		return log.Info()
	case services.LevelWarn:
		return log.Warn()
	case services.LevelError:
		return log.Error().Stack()
	default:
		return log.Info()
	}
}
