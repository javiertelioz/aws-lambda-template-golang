package services

import "context"

// Level represents the severity level of a log message.
type Level int

const (
	// LevelTrace is the most verbose level, used for detailed diagnostic information.
	LevelTrace Level = iota
	// LevelDebug is used for detailed information useful for debugging.
	LevelDebug
	// LevelInfo is used for general informational messages.
	LevelInfo
	// LevelWarn is used for potentially harmful situations.
	LevelWarn
	// LevelError is used for error events that might still allow the application to continue.
	LevelError
)

// String returns the string representation of the log level.
func (l Level) String() string {
	switch l {
	case LevelTrace:
		return "trace"
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	default:
		return "info"
	}
}

// Field represents a structured logging field with a key-value pair.
type Field struct {
	Key   string
	Value interface{}
}

// Logger defines a minimal interface for structured logging operations.
// This interface follows the principle "the bigger the interface, the weaker the abstraction".
//
// The single Log method accepts a context for propagating request IDs and trace information,
// a severity level, a message, and optional structured fields.
//
// Example usage:
//
//	logger.Log(ctx, services.LevelInfo, "User logged in",
//	    services.Field{Key: "user_id", Value: 123},
//	    services.Field{Key: "ip", Value: "192.168.1.1"},
//	)
//
// Benefits of a small interface:
//   - Easy to mock (only 1 method instead of 5)
//   - Extensible without breaking changes
//   - Supports structured logging natively
//   - Follows Go idioms (io.Reader, io.Writer pattern)
//   - Reduces code duplication in implementations
type Logger interface {
	Log(ctx context.Context, level Level, msg string, fields ...Field)
}
