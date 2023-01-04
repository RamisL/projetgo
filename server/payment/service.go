package payment

type Service interface {
	CreatePayment(input InputPayment) (Payment, error)
	GetAllPayment() ([]Payment, error)
	GetByIdPayment(id int) (Payment, error)
	UpdatePayment(id int, input InputPayment) (Payment, error)
	DeletePayment(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}
func (s *service) CreatePayment(input InputPayment) (Payment, error) {
	newPayment, err := s.repository.CreatePayment(input)
	if err != nil {
		return newPayment, err
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

func (s *service) DeletePayment(id int) error {
	err := s.repository.DeletePayment(id)
	if err != nil {
		return err
	}

	return nil
}
