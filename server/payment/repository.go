package payment

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/RamisL/server/product"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	CreatePayment(input InputPayment) (Payment, error)
	GetAllPayment() ([]Payment, error)
	GetByIdPayment(id int) (Payment, error)
	UpdatePayment(id int, input InputPayment) (Payment, error)
	DeletePayment(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) CreatePayment(input InputPayment) (Payment, error) {
	var payment Payment
	var product product.Product

	r.db.Select("price").Find(&product).Where("id = ?", input.ProductId).Scan(&product)
	payment.ProductId = input.ProductId
	payment.PricePaid = product.Price

	r.db.Where("id = ?", payment.ProductId).First(&product)

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
		r.db.Where("id = ?", payment.ProductId).First(&product)
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

	r.db.Where("id = ?", payment.ProductId).First(&product)
	payment.PricePaid = product.Price
	payment.ProductId = product.ID
	payment.Product = product

	return payment, nil
}

func (r *repository) UpdatePayment(id int, input InputPayment) (Payment, error) {
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
	var product product.Product

	db.Where("id = ?", input.ProductId).First(&product)
	payment.Product = product

	payment.ProductId = input.ProductId
	payment.PricePaid = product.Price

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
