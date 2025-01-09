package main

import (
	"fmt"
	"net/http"

	"github.com/build-a-thing-and-show/message-bus/endpoint"
	"github.com/build-a-thing-and-show/message-bus/service"
	"github.com/build-a-thing-and-show/message-bus/transport"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	// Initialize the service
	svc := service.NewHelloWorldService()

	// Create the greet endpoint
	greetEndpoint := endpoint.MakeGreetEndpoint(svc)

	// Create the HTTP handler
	handler := httptransport.NewServer(
		greetEndpoint,
		transport.DecodeGreetRequest,
		transport.EncodeResponse,
	)

	// Start the HTTP server
	http.Handle("/greet", handler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started at http://localhost:8080")
}
