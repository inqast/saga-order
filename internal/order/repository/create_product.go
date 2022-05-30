package repository

import (
	"context"

	"github.com/inqast/saga-order/internal/models"
)

func (r *repository) createProduct(ctx context.Context, orderId int, product *models.Product) (err error) {

	const query = `
		insert into products (
			order_id,
			product_id,
			count
		) VALUES (
			$1, $2, $3
		)
	`

	_, err = r.pool.Exec(ctx, query,
		orderId,
		product.Id,
		product.Count,
	)

	return
}
