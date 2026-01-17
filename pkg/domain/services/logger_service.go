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
// This method implements the fmt.Stringer interface, allowing Level to be used
// directly in formatted output, logging, and debugging.
//
// Example:
//
//	level := LevelInfo
//	fmt.Println(level.String())  // Output: "info"
//	fmt.Printf("Log level: %s", level)  // Uses String() automatically
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
// Context Usage:
// The context parameter is used to extract metadata for log correlation in distributed systems.
// Common context keys that implementations may extract:
//   - "request_id": Request identifier for correlating logs within a single request
//   - "trace_id": Distributed trace ID (AWS X-Ray, OpenTelemetry)
//   - "correlation_id": Correlation ID for tracking across multiple services
//   - "user_id": Authenticated user identifier
//
// Example usage:
//
//	// Basic logging
//	logger.Log(ctx, services.LevelInfo, "User logged in",
//	    services.Field{Key: "user_id", Value: 123},
//	    services.Field{Key: "ip", Value: "192.168.1.1"},
//	)
//
//	// With context values for tracing
//	ctx = context.WithValue(ctx, "request_id", "req-abc-123")
//	ctx = context.WithValue(ctx, "trace_id", "trace-xyz-789")
//	logger.Log(ctx, services.LevelInfo, "Processing order",
//	    services.Field{Key: "order_id", Value: 456},
//	)
//
// Benefits of a small interface:
//   - Easy to mock (only 1 method instead of 5)
//   - Extensible without breaking changes
//   - Supports structured logging natively
//   - Context propagation for distributed tracing
//   - Follows Go idioms (io.Reader, io.Writer pattern)
//   - Reduces code duplication in implementations
type Logger interface {
	// Log writes a log message at the specified level with optional structured fields.
	// The context can contain trace IDs, request IDs, or other metadata that will be
	// automatically extracted and added to the log entry for correlation.
	Log(ctx context.Context, level Level, msg string, fields ...Field)
}
