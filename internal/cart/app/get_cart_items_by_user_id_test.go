package app

import (
	"context"
	"github.com/inqast/saga-order/internal/cart/repository"
	"github.com/inqast/saga-order/internal/models"
	pb "github.com/inqast/saga-order/pkg/api/cart"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetCartItemsByUserId(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testUserID := 2

	testData := []*models.CartItem{
		{
			UserId:    2,
			ProductId: 2,
			Count:     5,
		},
		{
			UserId:    2,
			ProductId: 1,
			Count:     10,
		},
	}
	ctx := context.Background()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetCartItemsByUserIdMock.Return(testData, nil)
	mockRepo.GetCartItemsByUserIdMock.Expect(ctx, testUserID)
	svc := New(mockRepo)

	resp, err := svc.GetCartItemsByUserId(ctx, &pb.ID{Id: int64(testUserID)})

	assert.Nil(t, err)
	for i, cartItem := range resp.CartItems {
		testItem := testData[i]
		assert.Equal(t, cartItem.UserId, int64(testItem.UserId))
		assert.Equal(t, cartItem.ProductId, int64(testItem.ProductId))
		assert.Equal(t, cartItem.Count, int64(testItem.Count))
	}
}

func TestGetCartItemsNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.GetCartItemsByUserIdMock.Return([]*models.CartItem{}, repository.ErrNotFound)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.GetCartItemsByUserId(ctx, &pb.ID{Id: 1})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
