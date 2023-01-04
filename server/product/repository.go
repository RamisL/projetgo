package product

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	CreateProduct(product Product) (Product, error)
	GetAllProduct() ([]Product, error)
	GetByIdProduct(id int) (Product, error)
	UpdateProduct(id int, input InputProduct) (Product, error)
	DeleteProduct(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateProduct(product Product) (Product, error) {

	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
func (r *repository) GetAllProduct() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
func (r *repository) GetByIdProduct(id int) (Product, error) {
	var product Product

	err := r.db.Where(&Product{ID: id}).First(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) UpdateProduct(id int, input InputProduct) (Product, error) {
	product, err := r.GetByIdProduct(id)
	if err != nil {
		return product, err
	}

	product.Name = input.Name
	product.Price = input.Price

	err = r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
func (r *repository) DeleteProduct(id int) error {
	product := &Product{ID: id}
	px := r.db.Delete(product)
	if px.Error != nil {
		return px.Error
	}

	if px.RowsAffected == 0 {
		return errors.New("Product not found")
	}

	return nil
}
