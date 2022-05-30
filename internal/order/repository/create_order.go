package repository

import (
	"context"

	"github.com/inqast/saga-order/internal/models"
)

func (r *repository) CreateOrder(ctx context.Context, order *models.Order) (ID int, err error) {

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return
	}
	const query = `
		insert into orders (
			user_id,
			status
		) VALUES (
			$1, $2
		) returning id
	`

	err = r.pool.QueryRow(ctx, query,
		order.UserId,
		order.Status,
	).Scan(&ID)
	if err != nil {
		return
	}

	for _, product := range order.Products {
		err = r.createProduct(ctx, ID, product)
		if err != nil {
			_ = tx.Rollback(ctx)
			return
		}
	}

	err = tx.Commit(ctx)

	return
}
