package app

import "github.com/Shopify/sarama"

type Producer interface {
	SendMessage(*sarama.ProducerMessage) (partition int32, offset int64, err error)
}
