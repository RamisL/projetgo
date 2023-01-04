package payment

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/RamisL/server/product"
	product2 "github.com/RamisL/server/product"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	CreatePayment(payment Payment) (Payment, error)
	GetAllPayment() ([]Payment, error)
	GetByIdPayment(id int) (Payment, error)
	UpdatePayment(id int, input InputUpdatePayment) (Payment, error)
	DeletePayment(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) CreatePayment(payment Payment) (Payment, error) {

	var product product.Product
	r.db.Where("id = ?", payment.ProductID).First(&product)

	payment.Product = product

	err := r.db.Create(&payment).Error
	if err != nil {
		return payment, err
	}
	return payment, nil
}
func (r *repository) GetAllPayment() ([]Payment, error) {
	var payments []Payment
	err := r.db.Find(&payments).Error
	if err != nil {
		return nil, err
	}
	for key, payment := range payments {
		var product product.Product
		r.db.Where("id = ?", payment.ProductID).First(&product)
		fmt.Println(product)
		payments[key].Product = product

	}

	return payments, nil
}
func (r *repository) GetByIdPayment(id int) (Payment, error) {
	var payment Payment

	var product product.Product
	err := r.db.Where(&Payment{ID: id}).First(&payment).Error
	if err != nil {
		return payment, err
	}

	r.db.Where("id = ?", payment.ProductID).First(&product)
	payment.PricePaid = product.Price
	payment.ProductID = product.ID
	payment.Product = product

	return payment, nil
}

func (r *repository) UpdatePayment(id int, input InputUpdatePayment) (Payment, error) {
	payment, err := r.GetByIdPayment(id)
	if err != nil {
		return payment, err
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "user:password@tcp(127.0.0.1:3306)/projectgo?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	var product product2.Product

	db.Where("id = ?", payment.ProductID).First(&product)
	payment.Product = product

	payment.PricePaid = input.Price
	payment.Product.Name = input.Name
	payment.Product.Price = input.Price
	product.Name = input.Name
	product.Price = input.Price

	r.db.Save(&product)
	err = r.db.Save(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (r *repository) DeletePayment(id int) error {
	payment := &Payment{ID: id}
	px := r.db.Delete(payment)
	if px.Error != nil {
		return px.Error
	}

	if px.RowsAffected == 0 {
		return errors.New("Payment not found")
	}

	return nil
}
