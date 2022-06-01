package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/inqast/saga-order/internal/config"
	"github.com/inqast/saga-order/internal/models"
	"log"
	"time"
)

type NewOrderHandler struct {
	consumerGroup sarama.ConsumerGroup
	repo          Repository
	producer      Producer
}

func New(cfg *sarama.Config, repo Repository, producer Producer) (*NewOrderHandler, error) {
	reset, err := sarama.NewConsumerGroup([]string{"kafka:9092"}, "newOrder", cfg)
	if err != nil {
		return nil, err
	}

	rHandler := &NewOrderHandler{
		consumerGroup: reset,
		repo:          repo,
		producer:      producer,
	}

	return rHandler, nil
}

func (h *NewOrderHandler) StartConsuming(ctx context.Context, cfg config.KafkaConfig) {
	go func() {
		for {
			err := h.consumerGroup.Consume(ctx, []string{cfg.Topics.NewOrder}, h)
			if err != nil {
				log.Printf("reset consumer error: %v", err)
				time.Sleep(time.Second * 5)
			}
		}
	}()
}

func (h *NewOrderHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *NewOrderHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *NewOrderHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := context.Background()
	for msg := range claim.Messages() {
		var order models.Order
		err := json.Unmarshal(msg.Value, &order)
		if err != nil {
			log.Printf("new order data %v: %v", string(msg.Value), err)
			continue
		}

		for _, product := range order.Products {
			if err = h.repo.CreateReserve(ctx, &models.Reserve{
				OrderId:   order.Id,
				ProductId: product.Id,
				Count:     product.Count,
			}); err != nil {
				log.Printf("bad order: %v", order.Id)
				_ = h.repo.DeleteReservesByOrderId(ctx, order.UserId)
				break
			}
		}
		if err != nil {
			log.Printf("Order %v broken", order.Id)
			_ = h.sendResetOrder(&order)
			continue
		}

		err = h.sendNewReserves(&order)

		log.Printf("Order %v reserved", order.Id)
	}
	return nil
}

func (h *NewOrderHandler) sendNewReserves(order *models.Order) error {
	cartItems := make([]*models.CartItem, 0)
	for _, product := range order.Products {
		cartItems = append(cartItems, &models.CartItem{
			UserId:    order.UserId,
			ProductId: product.Id,
		})
	}

	encodedItems, err := json.Marshal(cartItems)
	if err != nil {
		log.Printf("encoding error: %v", err)
		return err
	}

	par, off, err := h.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "new_reserves",
		Key:   sarama.StringEncoder(fmt.Sprintf("%v", order.Id)),
		Value: sarama.ByteEncoder(encodedItems),
	})
	log.Printf("order %v -> %v; %v", par, off, err)
	if err != nil {
		log.Printf("sending data error: %v", err)
		return err
	}

	return err
}

func (h *NewOrderHandler) sendResetOrder(order *models.Order) error {
	encodedOrder, err := json.Marshal(order)
	if err != nil {
		log.Printf("encoding error: %v", err)
		return err
	}

	par, off, err := h.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "reset_orders",
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
