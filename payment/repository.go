package payment

import "gorm.io/gorm"

type Repository interface {
	StoreProduct(payment Payment) (Payment, error)
	ListAllProduct() ([]Payment, error)
	ShowProduct(id string) (Payment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
