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

func TestReadOrder(t *testing.T) {

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
	ctx := context.Background()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.ReadOrderMock.Return(&models.Order{
		Id:       testOrderId,
		UserId:   testUserId,
		Status:   testStatus,
		Products: testProducts,
	}, nil)
	mockRepo.ReadOrderMock.Expect(ctx, testOrderId)
	svc := New(mockRepo)

	order, err := svc.ReadOrder(ctx, &pb.ID{
		Id: int64(testOrderId),
	})

	assert.Nil(t, err)
	assert.Equal(t, order.Id, int64(testOrderId))
	assert.Equal(t, order.UserId, int64(testUserId))
	assert.Equal(t, order.Status, pb.Status(testStatus))
	for i, product := range order.Products {
		assert.Equal(t, product.Id, int64(testProducts[i].Id))
		assert.Equal(t, product.Count, int64(testProducts[i].Count))
	}
}

func TestReadOrderNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.ReadOrderMock.Return(&models.Order{}, repository.ErrNotFound)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.ReadOrder(ctx, &pb.ID{})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
