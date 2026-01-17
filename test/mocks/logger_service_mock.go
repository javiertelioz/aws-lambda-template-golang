package service

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/javiertelioz/aws-lambda-golang/pkg/domain/services"
)

// MockLogger is a mock implementation of the Logger interface for testing.
type MockLogger struct {
	mock.Mock
}

// Log mocks the Log method of the Logger interface.
func (m *MockLogger) Log(ctx context.Context, level services.Level, msg string, fields ...services.Field) {
	m.Called(ctx, level, msg, fields)
}
