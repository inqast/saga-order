package app

import (
	"context"
	"errors"

	"github.com/inqast/saga-order/internal/cart/repository"
	pb "github.com/inqast/saga-order/pkg/api/cart"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *tserver) ReadCartItem(ctx context.Context, req *pb.CartItemRequest) (*pb.CartItem, error) {

	cartItem, err := t.repo.ReadCartItem(ctx, int(req.UserId), int(req.ProductId))
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.CartItem{
		UserId:    int64(cartItem.UserId),
		ProductId: int64(cartItem.ProductId),
		Count:     int64(cartItem.Count),
	}, err
}
