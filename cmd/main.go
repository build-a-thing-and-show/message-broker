package main

import (
	"log"
	"net/http"

	"github.com/build-a-thing-and-show/message-broker/internal/router"
	"github.com/gorilla/handlers"
)

func main() {
	// Initialize router and configure routes
	r := router.NewRouter()

	// Set up CORS options
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Replace with your UI's origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	// Start HTTP server on port 10000
	log.Println("Starting server on :10000")
	if err := http.ListenAndServe(":10000", corsOptions(r)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
