package product

import (
	"errors"

	"github.com/mceciabate/web-server/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) ([]domain.Product, error)
	Create(p domain.Product) (domain.Product, error)
	Update(id int, p domain.Product) (domain.Product, error)
	Delete(id int) error
	Buy(code string, quantity int) error
	GetByCodeValue(code string) (domain.Product, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// GetAll devuelve todos los productos
func (s *service) GetAll() ([]domain.Product, error) {
	l := s.r.GetAll()
	return l, nil
}

// GetByID busca un producto por su id
func (s *service) GetByID(id int) (domain.Product, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

// SearchPriceGt busca productos por precio mayor que el precio dado
func (s *service) SearchPriceGt(price float64) ([]domain.Product, error) {
	l := s.r.SearchPriceGt(price)
	if len(l) == 0 {
		return []domain.Product{}, errors.New("no products found")
	}
	return l, nil
}

// Create agrega un nuevo producto
func (s *service) Create(p domain.Product) (domain.Product, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *service) Update(id int, p domain.Product) (domain.Product, error) {
	_, err := s.r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}

	p.Id = id
	err = s.r.Update(p)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

// Delete elimina un producto
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Buy compra un product
func (s *service) Buy(code string, quantity int) error {
	err := s.r.Buy(code, quantity)
	if err != nil {
		return err
	}
	return nil
}

// Devuelve un producto por code_value
func (s *service) GetByCodeValue(code string) (domain.Product, error) {
	p, e := s.r.GetByCodeValue(code)
	if e != nil {
		return domain.Product{}, e
	}
	return p, nil
}
