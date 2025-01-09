package service

import "context"

// HelloWorldService defines the service interface
type HelloWorldService interface {
	Greet(ctx context.Context, name string) (string, error)
}

// service implements HelloWorldService
type service struct{}

// NewHelloWorldService creates a new HelloWorldService instance
func NewHelloWorldService() HelloWorldService {
	return service{}
}

// Greet implements the Greet method
func (s service) Greet(ctx context.Context, name string) (string, error) {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!", nil
}
