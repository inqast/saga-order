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

func TestReadCartItem(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testUserId := 1
	testProductId := 2
	testCount := 3
	ctx := context.Background()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.ReadCartItemMock.Return(&models.CartItem{
		UserId:    testUserId,
		ProductId: testProductId,
		Count:     testCount,
	}, nil)
	mockRepo.ReadCartItemMock.Expect(ctx, testUserId, testProductId)
	svc := New(mockRepo)

	cartItem, err := svc.ReadCartItem(ctx, &pb.CartItemRequest{
		UserId:    int64(testUserId),
		ProductId: int64(testProductId),
	})

	assert.Nil(t, err)
	assert.Equal(t, cartItem.UserId, int64(testUserId))
	assert.Equal(t, cartItem.ProductId, int64(testProductId))
	assert.Equal(t, cartItem.Count, int64(testCount))
}

func TestReadCartItemNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.ReadCartItemMock.Return(&models.CartItem{}, repository.ErrNotFound)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.ReadCartItem(ctx, &pb.CartItemRequest{})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
