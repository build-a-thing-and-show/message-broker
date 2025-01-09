package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// GreetRequest represents the request payload
type GreetRequest struct {
	Name string
}

// GreetResponse represents the response payload
type GreetResponse struct {
	Message string `json:"message"`
	Err     string `json:"error,omitempty"`
}

// MakeGreetEndpoint wraps the Greet method in an endpoint
func MakeGreetEndpoint(svc HelloWorldService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GreetRequest) // Type assertion, in case of assertion failure, it should be caught. Furthur improvement
		message, err := svc.Greet(ctx, req.Name)
		if err != nil {
			return GreetResponse{Message: "", Err: err.Error()}, nil
		}
		return GreetResponse{Message: message, Err: ""}, nil
	}
}
