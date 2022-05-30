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

func (t *tserver) DeleteOrder(ctx context.Context, req *pb.ID) (*emptypb.Empty, error) {

	err := t.repo.DeleteOrder(ctx, int(req.Id))
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &emptypb.Empty{}, err
}
