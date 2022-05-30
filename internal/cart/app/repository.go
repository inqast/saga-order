package app

import (
	"context"
	"github.com/inqast/saga-order/internal/models"
)

type Repository interface {
	CreateCartItem(context.Context, *models.CartItem) error
	ReadCartItem(context.Context, int, int) (*models.CartItem, error)
	UpdateCartItem(context.Context, *models.CartItem) error
	DeleteCartItem(context.Context, int, int) error
	GetCartItemsByUserId(context.Context, int) ([]*models.CartItem, error)
}
