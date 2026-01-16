package hello

import (
	"fmt"
)

// SayHelloUseCase generates a personalized greeting message.
// It accepts a name and returns a formatted string with "Hello {name}!".
// If an empty name is provided, the caller should handle the default value.
//
// Example:
//
//	SayHelloUseCase("John") // returns "Hello John!"
func SayHelloUseCase(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
