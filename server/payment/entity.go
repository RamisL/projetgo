package payment

import (
	"github.com/RamisL/server/product"
	"time"
)

type Payment struct {
	ID        int             `json:"id" gorm:"primaryKey;autoIncrement:true"`
    ProductID int `gorm:"foreignkey:ProductID"`
	Product   product.Product `gorm:"foreignKey:ID;references:ProductID;constraint:OnDelete:SET NULL;"`
	PricePaid string          `json:"price_paid"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
