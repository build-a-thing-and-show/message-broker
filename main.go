package main

import (
	"log"
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
	log.Println("Starting HTTP server on :10000")
	log.Fatal(http.ListenAndServe(":10000", nil))
}
