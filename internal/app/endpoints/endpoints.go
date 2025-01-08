package endpoints

import (
	"context"

	"github.com/build-a-thing-and-show/message-bus/internal/app/services"

	"github.com/go-kit/kit/endpoint"
)

// HelloWorldRequest defines the input for the endpoint.
type HelloWorldRequest struct {
	Name string `json:"name"`
}

// HelloWorldResponse defines the output for the endpoint.
type HelloWorldResponse struct {
	Message string `json:"message"`
}

// MakeHelloWorldEndpoint creates the endpoint.
func MakeHelloWorldEndpoint(svc services.HelloWorldService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HelloWorldRequest) // Casting the request any type to HelloWorldRequest type
		message := svc.SayHello(req.Name)
		return HelloWorldResponse{Message: message}, nil
	}
}
