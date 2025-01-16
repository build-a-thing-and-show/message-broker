package service

import (
	"log"

	"github.com/build-a-thing-and-show/message-broker/internal/model"

	"github.com/build-a-thing-and-show/message-broker/internal/client"
)

// ProcessReceivedMessage processes the message received by the broker
func ProcessReceivedMessage(msg model.Message) error {
	// Here, you could implement custom logic, like storing the message, validating it, etc.
	log.Printf("Received message: %v", msg)

	// For now, just returning nil assuming no processing error
	return nil
}

// SendMessageToService sends the message to another microservice via an HTTP request
func SendMessageToService(msg model.Message) error {
	// Assuming target service URL is defined here, you may parameterize it
	targetServiceURL := "http://localhost:20000/api/receive"

	// Use the HTTP client to send the message to the target service
	if err := client.SendHTTPPostRequest(targetServiceURL, msg); err != nil {
		return err
	}

	log.Printf("Sent message: %v", msg)
	return nil
}
