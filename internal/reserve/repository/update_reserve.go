package repository

import (
	"context"
	"github.com/inqast/saga-order/internal/models"
)

func (r *repository) UpdateReserve(ctx context.Context, reserve *models.Reserve) (err error) {

	const query = `
		update reserve
		set count = $3
		where order_id = $1 and product_id = $2;
	`

	cmd, err := r.pool.Exec(ctx, query,
		reserve.OrderId,
		reserve.ProductId,
		reserve.Count,
	)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return
}
