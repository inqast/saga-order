package handler

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/inqast/saga-order/internal/config"
	"github.com/inqast/saga-order/internal/models"
	"log"
	"time"
)

type ResetHandler struct {
	consumerGroup sarama.ConsumerGroup
	repo          Repository
}

func New(cfg *sarama.Config, repo Repository) (*ResetHandler, error) {
	reset, err := sarama.NewConsumerGroup([]string{"kafka:9092"}, "resetOrder", cfg)
	if err != nil {
		return nil, err
	}

	rHandler := &ResetHandler{
		consumerGroup: reset,
		repo:          repo,
	}

	return rHandler, nil
}

func (r *ResetHandler) StartConsuming(ctx context.Context, cfg config.KafkaConfig) {
	go func() {
		for {
			err := r.consumerGroup.Consume(ctx, []string{cfg.Topics.ResetOrder}, r)
			if err != nil {
				log.Printf("reset consumer error: %v", err)
				time.Sleep(time.Second * 5)
			}
		}
	}()
}

func (r *ResetHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (r *ResetHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (r *ResetHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var order models.Order
		err := json.Unmarshal(msg.Value, &order)
		if err != nil {
			log.Printf("reset data %v: %v", string(msg.Value), err)
			continue
		}

		if err := r.repo.DeleteOrder(context.Background(), order.Id); err != nil {
			log.Printf("bad order: %v", order.Id)
			continue
		}
		log.Printf("Order %v deleted", order.Id)
	}
	return nil
}
