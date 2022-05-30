package app

import (
	"context"
	"errors"

	"github.com/inqast/saga-order/internal/cart/repository"
	pb "github.com/inqast/saga-order/pkg/api/cart"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *tserver) DeleteCartItem(ctx context.Context, req *pb.CartItemRequest) (*emptypb.Empty, error) {

	err := t.repo.DeleteCartItem(ctx, int(req.UserId), int(req.ProductId))
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &emptypb.Empty{}, err
}
