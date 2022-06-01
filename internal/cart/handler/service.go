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

type CleanHandler struct {
	consumerGroup sarama.ConsumerGroup
	repo          Repository
}

func New(cfg *sarama.Config, repo Repository) (*CleanHandler, error) {
	reset, err := sarama.NewConsumerGroup([]string{"kafka:9092"}, "newReserves", cfg)
	if err != nil {
		return nil, err
	}

	rHandler := &CleanHandler{
		consumerGroup: reset,
		repo:          repo,
	}

	return rHandler, nil
}

func (c *CleanHandler) StartConsuming(ctx context.Context, cfg config.KafkaConfig) {
	go func() {
		for {
			err := c.consumerGroup.Consume(ctx, []string{cfg.Topics.NewReserves}, c)
			if err != nil {
				log.Printf("reset consumer error: %v", err)
				time.Sleep(time.Second * 5)
			}
		}
	}()
}

func (c *CleanHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *CleanHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *CleanHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var cartItems []models.CartItem
		err := json.Unmarshal(msg.Value, &cartItems)
		if err != nil {
			log.Printf("reset data %v: %v", string(msg.Value), err)
			continue
		}

		for i := 0; i < 10; i++ {
			for _, cartItem := range cartItems {
				err = c.repo.DeleteCartItem(context.Background(), cartItem.UserId, cartItem.ProductId)
				if err != nil {
					log.Printf("bad item: user->%v cart->%v", cartItem.UserId, cartItem.ProductId)
					break
				}
			}
			if err == nil {
				break
			}
		}
	}
	return nil
}
