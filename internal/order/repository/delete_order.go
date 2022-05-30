package repository

import "context"

func (r *repository) DeleteOrder(ctx context.Context, Id int) (err error) {

	const query = `
		delete from orders
		where id = $1;
	`

	cmd, err := r.pool.Exec(ctx, query, Id)
	if cmd.RowsAffected() == 0 {
		err = ErrNotFound
		return
	}

	return
}
