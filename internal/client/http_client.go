package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/build-a-thing-and-show/message-broker/internal/model"
)

// SendHTTPPostRequest sends a POST request to the given URL with the provided message
func SendHTTPPostRequest(url string, msg model.Message) error {
	// Convert message to JSON
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// Send POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to send message: %v", resp.Status)
	}

	log.Printf("Message sent successfully to %s", url)
	return nil
}
