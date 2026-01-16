package service

import (
	"context"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockLoggerService struct {
	mock.Mock
}

func (ls *MockLoggerService) Trace(ctx context.Context, msg string) {
	fmt.Println(msg)
}

func (ls *MockLoggerService) Info(ctx context.Context, msg string) {
	fmt.Println(msg)
}

func (ls *MockLoggerService) Debug(ctx context.Context, msg string) {
	fmt.Println(msg)
}

func (ls *MockLoggerService) Warn(ctx context.Context, msg string) {
	fmt.Println(msg)
}

func (ls *MockLoggerService) Error(ctx context.Context, msg string) {
	fmt.Println(msg)
}
