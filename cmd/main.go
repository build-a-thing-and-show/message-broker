package main

import (
	"log"
	"net/http"

	"github.com/build-a-thing-and-show/message-broker/internal/router"
)

func main() {
	// Initialize router and configure routes
	r := router.NewRouter()

	// Start HTTP server on port 10000
	log.Println("Starting server on :10000")
	if err := http.ListenAndServe(":10000", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
