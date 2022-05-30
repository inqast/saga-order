package app

import (
	"context"
	"errors"
	"github.com/inqast/saga-order/internal/cart/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/inqast/saga-order/internal/models"
	pb "github.com/inqast/saga-order/pkg/api/cart"
)

func (t *tserver) CreateCartItem(ctx context.Context, req *pb.CartItem) (*emptypb.Empty, error) {

	var cartItem = models.CartItem{
		UserId:    int(req.UserId),
		ProductId: int(req.ProductId),
		Count:     int(req.Count),
	}

	err := t.repo.CreateCartItem(ctx, &cartItem)
	if errors.Is(err, repository.ErrAlreadyExists) {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	return &emptypb.Empty{}, err
}
