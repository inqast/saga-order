package repository

import (
	"context"
)

func (r *repository) UpdateOrder(ctx context.Context, id, status int) (err error) {

	const query = `
		update orders
		set status = $2
		where id = $1;
	`

	cmd, err := r.pool.Exec(ctx, query,
		id,
		status,
	)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return
}
