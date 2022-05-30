package repository

import "context"

func (r *repository) DeleteCartItem(ctx context.Context, userId, productId int) (err error) {

	const query = `
		delete from cart_items
		where user_id = $1 and product_id = $2;
	`

	cmd, err := r.pool.Exec(ctx, query, userId, productId)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return
}
