package handler

import (
	"context"
)

type Repository interface {
	DeleteOrder(context.Context, int) error
}
