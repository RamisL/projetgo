package product

type Service interface {
	CreateProduct(input InputProduct) (Product, error)
	GetAllProduct() ([]Product, error)
	GetByIdProduct(id int) (Product, error)
	UpdateProduct(id int, input InputProduct) (Product, error)
	DeleteProduct(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) CreateProduct(input InputProduct) (Product, error) {
	var product Product
	product.Name = input.Name
	product.Price = input.Price

	newProduct, err := s.repository.CreateProduct(product)
	if err != nil {
		return product, err
	}

	return newProduct, nil

}
func (s *service) GetAllProduct() ([]Product, error) {
	products, err := s.repository.GetAllProduct()
	if err != nil {
		return products, err
	}

	return products, nil
}
func (s *service) GetByIdProduct(id int) (Product, error) {
	product, err := s.repository.GetByIdProduct(id)
	if err != nil {
		return product, err
	}

	return product, nil
}
func (s *service) UpdateProduct(id int, input InputProduct) (Product, error) {
	uProduct, err := s.repository.UpdateProduct(id, input)
	if err != nil {
		return uProduct, err
	}

	return uProduct, nil
}

func (s *service) DeleteProduct(id int) error {
	err := s.repository.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
