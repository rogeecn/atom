package contracts

import "github.com/ThreeDotsLabs/watermill/message"

type EventHandler interface {
	Topic() string
	Channel() Channel
	PublishTo() (Channel, string)
	Handler(msg *message.Message) ([]*message.Message, error)
}

type Channel string

type EventPublisher interface {
	Topic() string
	Channel() Channel
	Marshal() ([]byte, error)
}
