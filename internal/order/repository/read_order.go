package repository

import (
	"context"
	"errors"

	"github.com/inqast/saga-order/internal/models"
	"github.com/jackc/pgx/v4"
)

func (r *repository) ReadOrder(ctx context.Context, id int) (*models.Order, error) {
	const query = `
		select id,
			user_id,
			status
		from orders
		where id = $1;
	`
	order := &models.Order{}
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&order.Id,
		&order.UserId,
		&order.Status,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}

	products, err := r.getProductsByOrderId(ctx, order.Id)
	if err != nil {
		return nil, ErrNotFound
	}
	order.Products = products

	return order, err
}
