package contracts

import "github.com/ThreeDotsLabs/watermill/message"

type EventHandler interface {
	Topic() string
	PublishToTopic() string
	Handler(msg *message.Message) ([]*message.Message, error)
}
