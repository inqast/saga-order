package repository

import (
	"context"
	"github.com/inqast/saga-order/internal/models"
)

func (r *repository) GetReservesByOrderId(ctx context.Context, orderID int) (reserves []*models.Reserve, err error) {
	const query = `
		select order_id,
			product_id,
			count
	  	from reserves
		where order_id = $1;
	`
	rows, err := r.pool.Query(ctx, query, orderID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var reserve models.Reserve
		if err = rows.Scan(
			&reserve.OrderId,
			&reserve.ProductId,
			&reserve.Count,
		); err != nil {
			return
		}

		reserves = append(reserves, &reserve)
	}

	return
}
