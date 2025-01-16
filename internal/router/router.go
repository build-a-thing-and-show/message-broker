package router

import (
	"github.com/build-a-thing-and-show/message-broker/internal/handler"

	"github.com/gorilla/mux"
)

// NewRouter returns a new HTTP router
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Define routes for the message broker
	r.HandleFunc("/api/message", handler.ReceiveMessage).Methods("POST")
	r.HandleFunc("/api/send", handler.SendMessage).Methods("POST")

	return r
}
