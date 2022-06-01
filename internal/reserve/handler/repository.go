package handler

import (
	"context"
	"github.com/inqast/saga-order/internal/models"
)

type Repository interface {
	CreateReserve(context.Context, *models.Reserve) error
	DeleteReservesByOrderId(context.Context, int) error
}
