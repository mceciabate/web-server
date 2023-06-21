package product

import (
	"errors"

	"github.com/mceciabate/web-server/internal/domain"
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
	list []domain.Product
}

// NewRepository crea un nuevo repositorio
func NewRepository(list []domain.Product) Repository {
	return &repository{list}
}

// GetAll devuelve todos los productos
func (r *repository) GetAll() []domain.Product {
	return r.list
}

// GetByID busca un producto por su id
func (r *repository) GetByID(id int) (domain.Product, error) {
	for _, product := range r.list {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")

}

func (r *repository) GetByCodeValue(code string) (domain.Product, error) {
	for _, p := range r.list {
		if p.CodeValue == code {
			return p, nil
		}
	}
	return domain.Product{}, errors.New("Product Not Found")
}

// SearchPriceGt busca productos por precio mayor o igual que el precio dado
func (r *repository) SearchPriceGt(price float64) []domain.Product {
	var products []domain.Product
	for _, product := range r.list {
		if product.Price > price {
			products = append(products, product)
		}
	}
	return products
}

// Create agrega un nuevo producto
func (r *repository) Create(p domain.Product) (domain.Product, error) {
	if !r.validateCodeValue(p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	p.Id = len(r.list) + 1
	r.list = append(r.list, p)
	return p, nil
}

// Actualizar un producto
func (r *repository) Update(p domain.Product) error {
	if !r.validateCodeValue(p.CodeValue) {
		return errors.New("code value already exist")
	}

	for i, prod := range r.list {
		if p.Id == prod.Id {
			r.list[i] = p
			return nil
		}
	}

	return errors.New("product not found")
}

// validateCodeValue valida que el codigo no exista en la lista de productos
func (r *repository) validateCodeValue(codeValue string) bool {
	for _, product := range r.list {
		if product.CodeValue == codeValue {
			return false
		}
	}
	return true
}

// Delete elimina un producto
func (r *repository) Delete(id int) error {
	for i, product := range r.list {
		if product.Id == id {
			r.list = append(r.list[:i], r.list[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}

// Setea la cantidad de prodcuto según la compra
func (r *repository) Buy(code string, quantity int) error {
	productForBuy, err := r.GetByCodeValue(code)
	if err != nil {
		return err
	}
	if productForBuy.Quantity < quantity {
		return errors.New("Cantidad insuficiente")
	}
	productForBuy.Quantity -= quantity
	return nil
}
