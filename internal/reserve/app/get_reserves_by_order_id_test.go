package app

import (
	"context"
	"testing"

	"github.com/inqast/saga-order/internal/models"
	"github.com/inqast/saga-order/internal/order/repository"
	pb "github.com/inqast/saga-order/pkg/api/reserve"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetReservesByOrderId(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testOrderID := 1

	testData := []*models.Reserve{
		{
			OrderId:   1,
			ProductId: 2,
			Count:     3,
		},
		{
			OrderId:   1,
			ProductId: 4,
			Count:     5,
		},
	}
	ctx := context.Background()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetReservesByOrderIdMock.Return(testData, nil)
	mockRepo.GetReservesByOrderIdMock.Expect(ctx, testOrderID)

	svc := New(mockRepo)

	resp, err := svc.GetReservesByOrderId(ctx, &pb.ID{Id: int64(testOrderID)})

	assert.Nil(t, err)
	for i, reserve := range resp.Reserves {
		testReserve := testData[i]
		assert.Equal(t, reserve.OrderId, int64(testReserve.OrderId))
		assert.Equal(t, reserve.ProductId, int64(testReserve.ProductId))
		assert.Equal(t, reserve.Count, int64(testReserve.Count))
	}
}

func TestGetReservesByOrderIdNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetReservesByOrderIdMock.Return([]*models.Reserve{}, repository.ErrNotFound)

	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.GetReservesByOrderId(ctx, &pb.ID{Id: 1})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
