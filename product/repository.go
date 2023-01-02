package product

import "gorm.io/gorm"

type Repository interface {
	StoreProduct(product Product) (Product, error)
	ListAllProduct() ([]Product, error)
	ShowProduct(id string) (Product, error)

	StorePayment(payment Payment) (Payment, error)
	ListAllPayment() ([]Payment, error)
	ShowPayment(id string) (Payment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
