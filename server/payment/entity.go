package payment

import (
	"github.com/RamisL/server/product"
	"time"
)

type Payment struct {
	ID        int             `json:"id" gorm:"primaryKey;autoIncrement:true"`
	ProductId int             `json:"product_id"`
	Product   product.Product `json:"product" gorm:"foreignKey:ProductId;references:ID;constraint:OnDelete:SET NULL;"`
	PricePaid string          `json:"price_paid"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
