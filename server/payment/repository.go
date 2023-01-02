package payment

import "gorm.io/gorm"

type Repository interface {
	CreatePayment(payment Payment) (Payment, error)
	UpdateProduct(id string) (Payment, error)
	GetAllProduct(id string) ([]Payment, error)
	GetByIdProduct(id string) (Payment, error)
	DeleteProduct(id string) (Payment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
