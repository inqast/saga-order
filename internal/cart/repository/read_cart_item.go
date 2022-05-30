package repository

import (
	"context"
	"errors"

	"github.com/inqast/saga-order/internal/models"
	"github.com/jackc/pgx/v4"
)

func (r *repository) ReadCartItem(ctx context.Context, userId, productId int) (*models.CartItem, error) {
	const query = `
		select user_id,
			product_id,
			count
		from cart_items
		where user_id = $1 and product_id = $2;
	`
	cartItem := &models.CartItem{}
	err := r.pool.QueryRow(ctx, query, userId, productId).Scan(
		&cartItem.UserId,
		&cartItem.ProductId,
		&cartItem.Count,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}

	return cartItem, err
}
