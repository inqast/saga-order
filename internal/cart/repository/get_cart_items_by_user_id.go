package repository

import (
	"context"

	"github.com/inqast/saga-order/internal/models"
)

func (r *repository) GetCartItemsByUserId(ctx context.Context, userID int) (cartItems []*models.CartItem, err error) {
	const query = `
		select user_id,
			product_id,
			count
	  	from cart_items
		where user_id = $1;
	`
	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cartItem models.CartItem
		if err = rows.Scan(
			&cartItem.UserId,
			&cartItem.ProductId,
			&cartItem.Count,
		); err != nil {
			return
		}

		cartItems = append(cartItems, &cartItem)
	}

	return
}
