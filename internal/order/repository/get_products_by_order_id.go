package repository

import (
	"context"

	"github.com/inqast/saga-order/internal/models"
)

func (r *repository) getProductsByOrderId(ctx context.Context, id int) (products []*models.Product, err error) {
	const query = `
		select product_id,
			count
	  	from products
		where order_id = $1;
	`
	rows, err := r.pool.Query(ctx, query, id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err = rows.Scan(
			&product.Id,
			&product.Count,
		); err != nil {
			return
		}

		products = append(products, &product)
	}

	return
}
