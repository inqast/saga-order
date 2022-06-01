package handler

import (
	"context"
)

type Repository interface {
	DeleteCartItem(context.Context, int, int) (err error)
}
