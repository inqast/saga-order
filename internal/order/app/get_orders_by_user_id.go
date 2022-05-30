package app

import (
	"context"
	"errors"

	"github.com/inqast/saga-order/internal/order/repository"
	pb "github.com/inqast/saga-order/pkg/api/order"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *tserver) GetOrdersByUserId(ctx context.Context, req *pb.ID) (*pb.GetOrdersResponse, error) {
	orders, err := t.repo.GetOrdersByUserId(ctx, int(req.Id))
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	externalOrders := make([]*pb.Order, 0)
	for _, order := range orders {
		externalOrder := &pb.Order{
			Id:     int64(order.Id),
			UserId: int64(order.UserId),
			Status: pb.Status(order.Status),
		}

		externalProducts := make([]*pb.Product, 0)
		for _, product := range order.Products {
			externalProducts = append(externalProducts, &pb.Product{
				Id:    int64(product.Id),
				Count: int64(product.Count),
			})
		}
		externalOrder.Products = externalProducts

		externalOrders = append(externalOrders, externalOrder)
	}

	return &pb.GetOrdersResponse{
		Orders: externalOrders,
	}, err
}
