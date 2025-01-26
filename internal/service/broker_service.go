package service

import (
	"log"

	"github.com/build-a-thing-and-show/message-broker/internal/model"

	"github.com/build-a-thing-and-show/message-broker/internal/client"
)

// ProcessReceivedMessage processes the message received by the broker
func ProcessReceivedMessage(msg model.Message) error {
	// Here, you could implement custom logic, like storing the message, validating it, etc.
	switch msg.ID {
	case "event-login-successful":
		// we receive a message from the login-service when the user is authenticated
		// we probably want the JWT token and userID to be stored somewhere that is accessible from all
		// microservices so they know the user is logged in if they make a request to them
		// We still have to figure out the expiry mechanism
	case "event-login-failed":
	case "event-user-registered":
	}
	log.Printf("Received message: %v", msg)
	SendMessageToService(msg)
	// For now, just returning nil assuming no processing error
	return nil
}

// SendMessageToService sends the message to another microservice via an HTTP request
func SendMessageToService(msg model.Message) error {
	// Assuming target service URL is defined here, you may parameterize it
	targetServiceURL := "http://localhost:20000/api/receive"
	coreUiURL := "http://localhost:5000/receive"

	// Use the HTTP client to send the message to the target service
	if err := client.SendHTTPPostRequest(targetServiceURL, msg); err != nil {
		return err
	}

	if err := client.SendHTTPPostRequest(coreUiURL, msg); err != nil {
		return err
	}

	log.Printf("Sent message: %v", msg)
	return nil
}
