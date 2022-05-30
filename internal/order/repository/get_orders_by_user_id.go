package repository

import (
	"context"

	"github.com/inqast/saga-order/internal/models"
)

func (r *repository) GetOrdersByUserId(ctx context.Context, userID int) (orders []*models.Order, err error) {
	const query = `
		select id,
			user_id,
			status
	  	from orders
		where user_id = $1;
	`
	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		if err = rows.Scan(
			&order.Id,
			&order.UserId,
			&order.Status,
		); err != nil {
			return
		}

		products, err := r.getProductsByOrderId(ctx, order.Id)
		if err != nil {
			return nil, err
		}
		order.Products = products

		orders = append(orders, &order)
	}

	return
}
