package app

import (
	"context"
	"errors"
	"github.com/inqast/saga-order/internal/cart/repository"
	pb "github.com/inqast/saga-order/pkg/api/cart"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *tserver) GetCartItemsByUserId(ctx context.Context, req *pb.ID) (*pb.GetCartItemsResponse, error) {

	cartItems, err := t.repo.GetCartItemsByUserId(ctx, int(req.Id))
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	externalCartItems := make([]*pb.CartItem, 0)
	for _, cartItem := range cartItems {
		externalCartItems = append(externalCartItems, &pb.CartItem{
			UserId:    int64(cartItem.UserId),
			ProductId: int64(cartItem.ProductId),
			Count:     int64(cartItem.Count),
		})
	}

	return &pb.GetCartItemsResponse{
		CartItems: externalCartItems,
	}, err
}
