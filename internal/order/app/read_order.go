package app

import (
	"context"
	"errors"
	"github.com/inqast/saga-order/internal/order/repository"
	pb "github.com/inqast/saga-order/pkg/api/order"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *tserver) ReadOrder(ctx context.Context, req *pb.ID) (*pb.Order, error) {

	order, err := t.repo.ReadOrder(ctx, int(req.Id))
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	products := make([]*pb.Product, 0)
	for _, product := range order.Products {
		products = append(products, &pb.Product{
			Id:    int64(product.Id),
			Count: int64(product.Count),
		})
	}

	return &pb.Order{
		Id:       int64(order.Id),
		UserId:   int64(order.UserId),
		Status:   pb.Status(order.Status),
		Products: products,
	}, err
}
