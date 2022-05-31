package app

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/inqast/saga-order/internal/models"
	"log"
)

func (t *tserver) sendToBroker(order *models.Order) error {
	encodedOrder, err := json.Marshal(order)
	if err != nil {
		log.Printf("encoding error: %v", err)
		return err
	}

	par, off, err := t.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "new_orders",
		Key:   sarama.StringEncoder(fmt.Sprintf("%v", order.Id)),
		Value: sarama.ByteEncoder(encodedOrder),
	})
	log.Printf("order %v -> %v; %v", par, off, err)
	if err != nil {
		log.Printf("sending data error: %v", err)
		return err
	}

	return err
}
