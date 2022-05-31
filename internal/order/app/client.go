package app

import (
	"context"
	"github.com/inqast/saga-order/internal/models"
)

type Client interface {
	GetCartItemsByUserId(ctx context.Context, id int) ([]*models.Product, error)
}
