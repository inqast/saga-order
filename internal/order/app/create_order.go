package app

import (
	"context"
	"github.com/inqast/saga-order/internal/models"
	pb "github.com/inqast/saga-order/pkg/api/order"
)

func (t *tserver) CreateOrder(ctx context.Context, req *pb.Order) (*pb.ID, error) {

	products, err := t.client.GetCartItemsByUserId(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}

	var order = &models.Order{
		Id:       int(req.Id),
		UserId:   int(req.UserId),
		Status:   int(req.Status),
		Products: products,
	}

	id, err := t.repo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	order.Id = id
	err = t.sendToBroker(order)
	if err != nil {
		t.repo.DeleteOrder(ctx, order.Id)
		return nil, err
	}

	return &pb.ID{Id: int64(id)}, err
}
