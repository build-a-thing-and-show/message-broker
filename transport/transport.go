// This is a basic hello world example. It is not applied in production anymore.
package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/build-a-thing-and-show/message-bus/endpoint"
)

// DecodeGreetRequest decodes an HTTP request into a GreetRequest
func DecodeGreetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.GreetRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

// EncodeResponse encodes the response as JSON
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
