package payment

import (
	product2 "github.com/RamisL/server/product"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Repository interface {
	CreatePayment(payment Payment) (Payment, error)
	GetAllPayment() ([]Payment, error)
	GetByIdPayment(id int) (Payment, error)
	UpdatePayment(id int, input InputPayment) (Payment, error)

	/*UpdateProduct(id string) (Payment, error)
	GetAllProduct(id string) ([]Payment, error)
	GetByIdProduct(id string) (Payment, error)
	DeleteProduct(id string) (Payment, error)*/
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) CreatePayment(payment Payment) (Payment, error) {

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

	return payments, nil
}
func (r *repository) GetByIdPayment(id int) (Payment, error) {
	var payment Payment

	err := r.db.Where(&Payment{ID: id}).First(&payment).Error
	if err != nil {
		return payment, err
	}

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
	var product product2.Product

	db.Select("price").Find(&product).Where("id = ?", input.ProductId).Scan(&product)

	payment.ProductId = input.ProductId
	payment.PricePaid = product.Price

	err = r.db.Save(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}
