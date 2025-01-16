package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/build-a-thing-and-show/message-broker/internal/model"
	"github.com/build-a-thing-and-show/message-broker/internal/service"
)

// ReceiveMessage handles incoming messages from other microservices
func ReceiveMessage(w http.ResponseWriter, r *http.Request) {
	var msg model.Message

	// Parse the message from the request body
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid message format", http.StatusBadRequest)
		return
	}

	// Call service to process the message
	if err := service.ProcessReceivedMessage(msg); err != nil {
		http.Error(w, fmt.Sprintf("Error processing message: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// SendMessage handles sending messages to other microservices
func SendMessage(w http.ResponseWriter, r *http.Request) {
	var msg model.Message

	// Parse the message from the request body
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid message format", http.StatusBadRequest)
		return
	}

	// Call service to send the message to the target microservice
	if err := service.SendMessageToService(msg); err != nil {
		http.Error(w, fmt.Sprintf("Error sending message: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
