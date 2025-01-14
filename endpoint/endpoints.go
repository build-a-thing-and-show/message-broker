// This is a basic hello world example. It is not applied in production anymore.
package endpoint

import (
	"context"

	"github.com/build-a-thing-and-show/message-bus/service"

	"github.com/go-kit/kit/endpoint"
)

// GreetRequest represents the input for the Greet endpoint
type GreetRequest struct {
	Name string `json:"name"`
}

// GreetResponse represents the output from the Greet endpoint
type GreetResponse struct {
	Message string `json:"message"`
	Err     string `json:"error,omitempty"`
}

// MakeGreetEndpoint creates an endpoint for the Greet method
func MakeGreetEndpoint(svc service.HelloWorldService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GreetRequest)
		message, err := svc.Greet(ctx, req.Name)
		if err != nil {
			return GreetResponse{Message: "", Err: err.Error()}, nil
		}
		return GreetResponse{Message: message, Err: ""}, nil
	}
}
