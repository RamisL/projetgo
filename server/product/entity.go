package product

import (
	"time"
)

type Product struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string    `json:"name" gorm:"unique"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
