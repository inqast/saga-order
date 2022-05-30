package app

import (
	"context"
	"testing"

	"github.com/inqast/saga-order/internal/cart/repository"
	pb "github.com/inqast/saga-order/pkg/api/cart"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeleteCartItem(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testUserId := 1
	testProductId := 2
	ctx := context.Background()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.DeleteCartItemMock.Return(nil)
	mockRepo.DeleteCartItemMock.Expect(ctx, testUserId, testProductId)
	svc := New(mockRepo)

	_, err := svc.DeleteCartItem(ctx, &pb.CartItemRequest{
		UserId:    int64(testUserId),
		ProductId: int64(testProductId),
	})

	assert.Nil(t, err)
}

func TestDeleteCartItemNotFound(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.DeleteCartItemMock.Return(repository.ErrNotFound)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.DeleteCartItem(ctx, &pb.CartItemRequest{})

	assert.Equal(t, err, status.Error(codes.NotFound, repository.ErrNotFound.Error()))
}
