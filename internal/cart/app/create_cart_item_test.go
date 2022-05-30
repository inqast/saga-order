package app

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/inqast/saga-order/internal/cart/repository"
	"github.com/inqast/saga-order/internal/models"
	pb "github.com/inqast/saga-order/pkg/api/cart"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestCreateCartItem(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	testUserId := 1
	testProductId := 2
	testCount := 3
	ctx := context.Background()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.CreateCartItemMock.Return(nil)
	mockRepo.CreateCartItemMock.Expect(ctx, &models.CartItem{
		UserId:    testUserId,
		ProductId: testProductId,
		Count:     testCount,
	})
	svc := New(mockRepo)

	_, err := svc.CreateCartItem(ctx, &pb.CartItem{
		UserId:    int64(testUserId),
		ProductId: int64(testProductId),
		Count:     int64(testCount),
	})

	assert.Nil(t, err)
}

func TestCreateCartItemAlreadyExist(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	mockRepo := NewRepositoryMock(mc)
	mockRepo.CreateCartItemMock.Return(repository.ErrAlreadyExists)
	svc := New(mockRepo)

	ctx := context.Background()
	_, err := svc.CreateCartItem(ctx, &pb.CartItem{})

	assert.Equal(t, err, status.Error(codes.AlreadyExists, repository.ErrAlreadyExists.Error()))
}
