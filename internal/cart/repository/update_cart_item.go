package repository

import (
	"context"

	"github.com/inqast/saga-order/internal/models"
)

func (r *repository) UpdateCartItem(ctx context.Context, cartItem *models.CartItem) (err error) {

	const query = `
		update cart_items
		set count = $3
		where user_id = $1 and product_id = $2;
	`

	cmd, err := r.pool.Exec(ctx, query,
		cartItem.UserId,
		cartItem.ProductId,
		cartItem.Count,
	)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return
}
