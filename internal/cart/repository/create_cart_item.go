package repository

import (
	"context"

	"github.com/inqast/saga-order/internal/models"
)

func (r *repository) CreateCartItem(ctx context.Context, cartItem *models.CartItem) (err error) {

	const query = `
		insert into cart_items (
			user_id,
			product_id,
			count
		) VALUES (
			$1, $2, $3
		)
	`

	cmd, err := r.pool.Exec(ctx, query,
		cartItem.UserId,
		cartItem.ProductId,
		cartItem.Count,
	)

	if cmd.RowsAffected() == 0 {
		err = ErrAlreadyExists
		return
	}

	return
}
