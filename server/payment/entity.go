package payment

import "time"

type Payment struct {
	ID        int       `json:"id"`
	productId int       `json:"product_id"`
	pricePaid string    `json:"price_paid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
