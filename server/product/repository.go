package product

import "gorm.io/gorm"

type Repository interface {
	CreateProduct(product Product) (Product, error)
	UpdateProduct(id string) (Product, error)
	GetAllProduct(id string) ([]Product, error)
	GetByIdProduct(id string) (Product, error)
	DeleteProduct(id string) (Product, error)
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
func (r *repository) ListAll() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
func (r *repository) Show(id string) (Task, error) {
	var task Task
	err := r.db.First(&id).Error
	if err != nil {
		return task, err
	}
	return task, nil
}
