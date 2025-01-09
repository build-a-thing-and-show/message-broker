package main

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// DecodeGreetRequest decodes an HTTP request into a GreetRequest
func DecodeGreetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GreetRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err
}

// EncodeResponse encodes the response as JSON
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	svc := service{}
	greetEndpoint := MakeGreetEndpoint(svc)

	handler := httptransport.NewServer(
		greetEndpoint,
		DecodeGreetRequest,
		EncodeResponse,
	)

	http.Handle("/greet", handler)
	http.ListenAndServe(":8080", nil)
}
