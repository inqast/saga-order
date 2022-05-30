package app

import (
	"context"
	"github.com/inqast/saga-order/internal/models"
	pb "github.com/inqast/saga-order/pkg/api/order"
)

func (t *tserver) CreateOrder(ctx context.Context, req *pb.Order) (*pb.ID, error) {

	products := make([]*models.Product, 0)
	for _, product := range req.Products {
		products = append(products, &models.Product{
			Id:    int(product.Id),
			Count: int(product.Count),
		})
	}

	var order = models.Order{
		Id:       int(req.Id),
		UserId:   int(req.UserId),
		Status:   int(req.Status),
		Products: products,
	}

	id, err := t.repo.CreateOrder(ctx, &order)

	return &pb.ID{Id: int64(id)}, err
}
