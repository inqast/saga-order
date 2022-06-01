package repository

import "context"

func (r *repository) DeleteReservesByOrderId(ctx context.Context, orderId int) (err error) {

	const query = `
		delete from reserves
		where order_id = $1;
	`

	cmd, err := r.pool.Exec(ctx, query, orderId)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return
}
