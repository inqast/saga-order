package app

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/inqast/saga-order/internal/models"
	pb "github.com/inqast/saga-order/pkg/api/order"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateOrder(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testOrderId := 1
	testUserId := 2
	testStatus := 0
	testExternalProducts := []*pb.Product{
		{
			Id:    3,
			Count: 4,
		},
		{
			Id:    5,
			Count: 6,
		},
	}
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
	ctx := context.Background()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.CreateOrderMock.Return(testOrderId, nil)
	mockRepo.CreateOrderMock.Expect(ctx, &models.Order{
		UserId:   testUserId,
		Status:   testStatus,
		Products: testProducts,
	})
	svc := New(mockRepo)

	_, err := svc.CreateOrder(ctx, &pb.Order{
		UserId:   int64(testUserId),
		Status:   pb.Status(testStatus),
		Products: testExternalProducts,
	})

	assert.Nil(t, err)
}
