package logger

import (
	"context"
	"fmt"
	"runtime"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/javiertelioz/aws-lambda-golang/pkg/domain/services"
)

// ZerologLogger is a concrete implementation of the LoggerService interface
// using the zerolog library for structured JSON logging.
// It includes automatic caller location tracking (file and line number)
// and stack trace marshaling for error logs.
type ZerologLogger struct{}

// NewLogger creates and configures a new ZerologLogger instance.
// It sets up zerolog with Unix timestamp format and stack trace support.
//
// Returns:
//   - A LoggerService implementation ready for use
//
// The logger outputs structured JSON logs to stdout with the following features:
//   - Unix timestamp format
//   - Caller location (file:line)
//   - Stack traces for error logs
func NewLogger() services.LoggerService {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	return &ZerologLogger{}
}

func (z *ZerologLogger) Trace(ctx context.Context, msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	log.Trace().
		Str("loc", fmt.Sprintf("%s:%d", file, line)).
		Msg(msg)
}

func (z *ZerologLogger) Debug(ctx context.Context, msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	log.Debug().
		Str("loc", fmt.Sprintf("%s:%d", file, line)).
		Msg(msg)
}

func (z *ZerologLogger) Info(ctx context.Context, msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	log.Info().
		Str("loc", fmt.Sprintf("%s:%d", file, line)).
		Msg(msg)
}

func (z *ZerologLogger) Warn(ctx context.Context, msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	log.Warn().
		Str("loc", fmt.Sprintf("%s:%d", file, line)).
		Msg(msg)
}

func (z *ZerologLogger) Error(ctx context.Context, msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	log.Error().
		Stack().
		Str("loc", fmt.Sprintf("%s:%d", file, line)).
		Msg(msg)
}
