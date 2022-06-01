package repository

import (
	"context"
	"github.com/inqast/saga-order/internal/models"
)

func (r *repository) CreateReserve(ctx context.Context, reserve *models.Reserve) error {

	const query = `
		insert into reserves (
			order_id,
			product_id,
			count
		) VALUES (
			$1, $2, $3
		)
	`

	cmd, err := r.pool.Exec(ctx, query,
		reserve.OrderId,
		reserve.ProductId,
		reserve.Count,
	)

	if cmd.RowsAffected() == 0 {
		err = ErrAlreadyExists
	}

	return err
}
