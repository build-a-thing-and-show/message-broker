package service

import "context"

type MessageBroker interface {
	Publish(ctx context.Context, msg *Message) error
	Subscribe(ctx context.Context, topic string) (<-chan *Message, error)
}
