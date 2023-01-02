package payment

import "time"

type Payment struct {
	ID        int       `json:"id"`
	productId int       `json:"productId"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
