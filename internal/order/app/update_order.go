package app

import (
	"context"
	"errors"

	"github.com/inqast/saga-order/internal/order/repository"
	pb "github.com/inqast/saga-order/pkg/api/order"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *tserver) UpdateOrder(ctx context.Context, req *pb.Order) (*emptypb.Empty, error) {

	err := t.repo.UpdateOrder(ctx, int(req.Id), int(req.Status))
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &emptypb.Empty{}, err
}
