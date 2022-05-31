package app

import (
	"context"
	"testing"

	"github.com/inqast/saga-order/internal/order/repository"
	pb "github.com/inqast/saga-order/pkg/api/order"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestUpdateOrder(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testOrderId := 1
	testUserId := 2
	testStatus := 0
	ctx := context.Background()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.UpdateOrderMock.Return(nil)
	mockRepo.UpdateOrderMock.Expect(ctx, testOrderId, testStatus)
	mockClient := NewClientMock(mc)
	mockProducer := NewProducerMock(mc)

	svc, err := New(mockRepo, mockClient, mockProducer)

	_, err = svc.UpdateOrder(ctx, &pb.Order{
		Id:     int64(testOrderId),
		UserId: int64(testUserId),
		Status: pb.Status(testStatus),
	})

	assert.Nil(t, err)
}

func TestUpdateOrderNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.UpdateOrderMock.Return(repository.ErrNotFound)
	mockClient := NewClientMock(mc)
	mockProducer := NewProducerMock(mc)

	svc, err := New(mockRepo, mockClient, mockProducer)

	ctx := context.Background()
	_, err = svc.UpdateOrder(ctx, &pb.Order{})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
