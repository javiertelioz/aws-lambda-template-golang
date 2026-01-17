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
func NewLogger() *ZerologLogger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return &ZerologLogger{}
}

func (z *ZerologLogger) Log(ctx context.Context, level services.Level, msg string, fields ...services.Field) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	event := z.getEventForLevel(level)
	event = event.Str("file", file).Int("line", line)
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
