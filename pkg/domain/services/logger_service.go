package services

import "context"

// LoggerService defines the interface for structured logging operations.
// All logging methods accept a context for propagating request IDs, trace information,
// and handling timeouts/cancellations.
//
// Implementations should output structured logs (e.g., JSON format) with:
//   - Timestamp
//   - Log level
//   - Message
//   - Caller location (file and line number)
//   - Any context-derived fields (request ID, trace ID, etc.)
//
// The interface provides five standard log levels from most verbose to most severe:
// Trace, Debug, Info, Warn, and Error.
type LoggerService interface {
	// Trace logs a message at TRACE level (most verbose).
	// Use for detailed diagnostic information during development.
	Trace(ctx context.Context, msg string)

	// Debug logs a message at DEBUG level.
	// Use for detailed information useful for debugging.
	Debug(ctx context.Context, msg string)

	// Info logs a message at INFO level.
	// Use for general informational messages about application flow.
	Info(ctx context.Context, msg string)

	// Warn logs a message at WARN level.
	// Use for potentially harmful situations that don't prevent operation.
	Warn(ctx context.Context, msg string)

	// Error logs a message at ERROR level.
	// Use for error events that might still allow the application to continue.
	Error(ctx context.Context, msg string)
}
