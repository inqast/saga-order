package app

import (
	"context"
	"github.com/inqast/saga-order/internal/models"
)

type Repository interface {
	CreateOrder(context.Context, *models.Order) (int, error)
	ReadOrder(context.Context, int) (*models.Order, error)
	UpdateOrder(context.Context, int, int) error
	DeleteOrder(context.Context, int) error
	GetOrdersByUserId(context.Context, int) ([]*models.Order, error)
}
