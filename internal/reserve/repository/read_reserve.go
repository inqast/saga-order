package repository

import (
	"context"
	"errors"

	"github.com/inqast/saga-order/internal/models"
	"github.com/jackc/pgx/v4"
)

func (r *repository) ReadReserve(ctx context.Context, id int) (*models.Reserve, error) {
	const query = `
		select order_id,
			product_id,
			count
		from reserves
		where order_id = $1;
	`
	reserve := &models.Reserve{}
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&reserve.OrderId,
		&reserve.ProductId,
		&reserve.Count,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}

	return reserve, err
}
