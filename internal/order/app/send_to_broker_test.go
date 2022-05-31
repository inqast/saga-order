package app

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gojuno/minimock/v3"
	"github.com/inqast/saga-order/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendToBroker(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testOrderId := 1
	testUserId := 2
	testStatus := 0
	testProducts := []*models.Product{
		{
			Id:    3,
			Count: 4,
		},
		{
			Id:    5,
			Count: 6,
		},
	}
	testEncodedOrder, _ := json.Marshal(&models.Order{
		Id:       testOrderId,
		UserId:   testUserId,
		Status:   testStatus,
		Products: testProducts,
	})

	mockRepo := NewRepositoryMock(mc)
	mockClient := NewClientMock(mc)
	mockProducer := NewProducerMock(mc)
	mockProducer.SendMessageMock.Return(0, 2, nil)
	mockProducer.SendMessageMock.Expect(&sarama.ProducerMessage{
		Topic: "new_orders",
		Key:   sarama.StringEncoder(fmt.Sprintf("%v", testOrderId)),
		Value: sarama.ByteEncoder(testEncodedOrder),
	})

	svc, err := New(mockRepo, mockClient, mockProducer)

	err = svc.sendToBroker(&models.Order{
		Id:       testOrderId,
		UserId:   testUserId,
		Status:   testStatus,
		Products: testProducts,
	})

	assert.Nil(t, err)
}
