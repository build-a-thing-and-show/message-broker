package model

// Message represents the structure of the message exchanged between services
type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}
