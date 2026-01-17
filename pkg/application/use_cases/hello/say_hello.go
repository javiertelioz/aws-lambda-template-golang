package hello

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	// MaxNameLength defines the maximum allowed length for a name
	MaxNameLength = 100
	// DefaultName is used when no name is provided
	DefaultName = "world"
)

// Validation errors
var (
	// ErrNameTooLong indicates that the name exceeds the maximum allowed length
	ErrNameTooLong = errors.New("name exceeds maximum length")

	// ErrInvalidCharacters indicates that the name contains invalid characters
	ErrInvalidCharacters = errors.New("name contains invalid characters")
)

// validNamePattern matches alphanumeric characters, spaces, hyphens, apostrophes, and common international characters
var validNamePattern = regexp.MustCompile(`^[a-zA-Z0-9\s\-'áéíóúÁÉÍÓÚñÑüÜ]+$`)

// SayHelloUseCase generates a personalized greeting message.
// It validates and sanitizes the input name according to business rules:
//   - Trims whitespace
//   - Uses default "world" if empty
//   - Validates length (max 100 characters)
//   - Validates allowed characters (alphanumeric, spaces, hyphens, apostrophes)
//
// Returns the greeting message and any validation error.
//
// Example:
//
//	SayHelloUseCase("John")           // returns "Hello John!", nil
//	SayHelloUseCase("")               // returns "Hello world!", nil
//	SayHelloUseCase("Very long...")   // returns "", ErrNameTooLong
//	SayHelloUseCase("<script>")       // returns "", ErrInvalidCharacters
func SayHelloUseCase(name string) (string, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return fmt.Sprintf("Hello %s!", DefaultName), nil
	}

	if len(name) > MaxNameLength {
		return "", ErrNameTooLong
	}

	if !validNamePattern.MatchString(name) {
		return "", ErrInvalidCharacters
	}

	return fmt.Sprintf("Hello %s!", name), nil
}
