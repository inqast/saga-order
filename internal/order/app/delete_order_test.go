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

func TestDeleteOrder(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testOrderId := 1
	ctx := context.Background()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.DeleteOrderMock.Return(nil)
	mockRepo.DeleteOrderMock.Expect(ctx, testOrderId)
	mockClient := NewClientMock(mc)
	mockProducer := NewProducerMock(mc)

	svc, err := New(mockRepo, mockClient, mockProducer)

	_, err = svc.DeleteOrder(ctx, &pb.ID{
		Id: int64(testOrderId),
	})

	assert.Nil(t, err)
}

func TestDeleteOrderNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.DeleteOrderMock.Return(repository.ErrNotFound)

	mockClient := NewClientMock(mc)
	mockProducer := NewProducerMock(mc)

	svc, err := New(mockRepo, mockClient, mockProducer)

	ctx := context.Background()
	_, err = svc.DeleteOrder(ctx, &pb.ID{})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
