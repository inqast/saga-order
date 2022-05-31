package cart

import (
	"context"
	"github.com/inqast/saga-order/internal/models"
	pb "github.com/inqast/saga-order/pkg/api/cart"
)

func (s *Service) GetCartItemsByUserId(ctx context.Context, id int) ([]*models.Product, error) {
	msg := pb.ID{
		Id: int64(id),
	}

	resp, err := s.grpcClient.GetCartItemsByUserId(ctx, &msg)
	if err != nil {
		return []*models.Product{}, err
	}

	products := make([]*models.Product, len(resp.CartItems))
	for i, product := range resp.CartItems {
		products[i] = &models.Product{
			Id:    int(product.ProductId),
			Count: int(product.Count),
		}
	}

	return products, err
}
