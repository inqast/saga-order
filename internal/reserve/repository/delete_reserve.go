package repository

import "context"

func (r *repository) DeleteReserve(ctx context.Context, orderId, productId int) (err error) {

	const query = `
		delete from reserves
		where order_id = $1 and product_id = $2;
	`

	cmd, err := r.pool.Exec(ctx, query, orderId, productId)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return
}
