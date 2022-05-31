package app

import (
	"context"
	"testing"

	"github.com/inqast/saga-order/internal/models"
	"github.com/inqast/saga-order/internal/order/repository"
	pb "github.com/inqast/saga-order/pkg/api/order"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetOrdersByUserId(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testUserID := 2

	testData := []*models.Order{
		{
			Id:     1,
			UserId: 2,
			Status: 3,
			Products: []*models.Product{
				{
					Id:    4,
					Count: 5,
				},
			},
		},
		{
			Id:     6,
			UserId: 2,
			Status: 8,
			Products: []*models.Product{
				{
					Id:    9,
					Count: 10,
				},
			},
		},
	}
	ctx := context.Background()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetOrdersByUserIdMock.Return(testData, nil)
	mockRepo.GetOrdersByUserIdMock.Expect(ctx, testUserID)
	mockClient := NewClientMock(mc)
	mockProducer := NewProducerMock(mc)

	svc, err := New(mockRepo, mockClient, mockProducer)

	resp, err := svc.GetOrdersByUserId(ctx, &pb.ID{Id: int64(testUserID)})

	assert.Nil(t, err)
	for i, order := range resp.Orders {
		testOrder := testData[i]
		assert.Equal(t, order.Id, int64(testOrder.Id))
		assert.Equal(t, order.UserId, int64(testOrder.UserId))
		assert.Equal(t, order.Status, pb.Status(testOrder.Status))
		for j, product := range order.Products {
			testProduct := testOrder.Products[j]
			assert.Equal(t, product.Id, int64(testProduct.Id))
			assert.Equal(t, product.Count, int64(testProduct.Count))
		}
	}
}

func TestGetOrdersByUserIdNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetOrdersByUserIdMock.Return([]*models.Order{}, repository.ErrNotFound)
	mockClient := NewClientMock(mc)
	mockProducer := NewProducerMock(mc)

	svc, err := New(mockRepo, mockClient, mockProducer)

	ctx := context.Background()
	_, err = svc.GetOrdersByUserId(ctx, &pb.ID{Id: 1})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
