package app

import (
	"context"
	"errors"
	"github.com/inqast/saga-order/internal/order/repository"
	pb "github.com/inqast/saga-order/pkg/api/reserve"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *tserver) GetReservesByOrderId(ctx context.Context, req *pb.ID) (*pb.GetReservesResponse, error) {
	reserves, err := t.repo.GetReservesByOrderId(ctx, int(req.Id))
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	externalReserves := make([]*pb.Reserve, 0)
	for _, reserve := range reserves {
		externalReserve := &pb.Reserve{
			OrderId:   int64(reserve.OrderId),
			ProductId: int64(reserve.ProductId),
			Count:     int64(reserve.Count),
		}

		externalReserves = append(externalReserves, externalReserve)
	}

	return &pb.GetReservesResponse{
		Reserves: externalReserves,
	}, err
}
