package main

import "context"

// Service interface defines the methods of our service
type HelloWorldService interface {
	Greet(ctx context.Context, name string) (string, error)
}

// service implements HelloWorldService
type service struct{}

func (s service) Greet(ctx context.Context, name string) (string, error) {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!", nil
}
