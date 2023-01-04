package payment

import (
	product2 "github.com/RamisL/server/product"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Service interface {
	CreatePayment(input InputPayment) (Payment, error)
	GetAllPayment() ([]Payment, error)
	GetByIdPayment(id int) (Payment, error)
	UpdatePayment(id int, input InputPayment) (Payment, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}
func (s *service) CreatePayment(input InputPayment) (Payment, error) {
	var payment Payment
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

	newPayment, err := s.repository.CreatePayment(payment)
	if err != nil {
		return payment, err
	}

	return newPayment, nil

}
func (s *service) GetAllPayment() ([]Payment, error) {
	payments, err := s.repository.GetAllPayment()
	if err != nil {
		return payments, err
	}

	return payments, nil
}
func (s *service) GetByIdPayment(id int) (Payment, error) {
	payment, err := s.repository.GetByIdPayment(id)
	if err != nil {
		return payment, err
	}

	return payment, nil
}
func (s *service) UpdatePayment(id int, input InputPayment) (Payment, error) {
	uPayment, err := s.repository.UpdatePayment(id, input)
	if err != nil {
		return uPayment, err
	}

	return uPayment, nil
}
