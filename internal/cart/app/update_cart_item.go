package app

import (
	"context"
	"errors"

	"github.com/inqast/saga-order/internal/cart/repository"
	"github.com/inqast/saga-order/internal/models"
	pb "github.com/inqast/saga-order/pkg/api/cart"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *tserver) UpdateCartItem(ctx context.Context, req *pb.CartItem) (*emptypb.Empty, error) {

	var cartItem = models.CartItem{
		UserId:    int(req.UserId),
		ProductId: int(req.ProductId),
		Count:     int(req.Count),
	}

	err := t.repo.UpdateCartItem(ctx, &cartItem)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &emptypb.Empty{}, err
}
