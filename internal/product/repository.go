package product

import (
	"errors"

	"github.com/mceciabate/web-server/internal/domain"
	"github.com/mceciabate/web-server/pkg/store"
)

type Repository interface {
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) []domain.Product
	Create(p domain.Product) (domain.Product, error)
	Update(p domain.Product) error
	Delete(id int) error
	GetByCodeValue(code string) (domain.Product, error)
	Buy(code string, quantity int) error
}

type repository struct {
	storage store.Store
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.Store) Repository {
	return &repository{storage}
}

// GetAll devuelve todos los productos
func (r *repository) GetAll() []domain.Product {
	products, err := r.storage.GetAll()
	if err != nil {
		return []domain.Product{}
	}
	return products
}

// GetByID busca un producto por su id
func (r *repository) GetByID(id int) (domain.Product, error) {
	product, err := r.storage.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil

}

func (r *repository) GetByCodeValue(code string) (domain.Product, error) {
	product, err := r.storage.GetByCodeValue(code)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

// SearchPriceGt busca productos por precio mayor o igual que el precio dado
func (r *repository) SearchPriceGt(price float64) []domain.Product {
	products := r.storage.SearchPriceGt(price)
	return products
}

// Create agrega un nuevo producto
func (r *repository) Create(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	err := r.storage.Create(p)
	if err != nil {
		return domain.Product{}, errors.New("error creating product")
	}
	return p, nil
}

// Actualizar un producto
func (r *repository) Update(p domain.Product) error {
	if !r.validateCodeValue(p.CodeValue) {
		errors.New("code value already exists")
	}
	err := r.storage.Update(p)
	if err != nil {
		errors.New("error updating product")
	}
	return nil
}

// validateCodeValue valida que el codigo no exista en la lista de productos
func (r *repository) validateCodeValue(codeValue string) bool {
	list, err := r.storage.GetAll()
	if err != nil {
		return false
	}
	for _, product := range list {
		if product.CodeValue == codeValue {
			return false
		}
	}
	return true
}

// Delete elimina un producto
func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Setea la cantidad de prodcuto según la compra
func (r *repository) Buy(code string, quantity int) error {
	err := r.storage.Buy(code, quantity)
	if err != nil {
		return err
	}
	return nil
}
