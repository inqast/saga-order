package app

import (
	"context"
	"github.com/inqast/saga-order/internal/models"
)

type Repository interface {
	GetReservesByOrderId(context.Context, int) ([]*models.Reserve, error)
}
