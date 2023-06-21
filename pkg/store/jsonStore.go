package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/mceciabate/web-server/internal/domain"
)

type Store interface {
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) []domain.Product
	Create(product domain.Product) error
	Update(product domain.Product) error
	Delete(id int) error
	GetByCodeValue(code string) (domain.Product, error)
	Buy(code string, quantity int) error
	saveProducts(products []domain.Product) error
	loadProducts() ([]domain.Product, error)
}

type jsonStore struct {
	pathToFile string
}

// loadProducts carga los productos desde un archivo json
func (s *jsonStore) loadProducts() ([]domain.Product, error) {
	var products []domain.Product
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// saveProducts guarda los productos en un archivo json
func (s *jsonStore) saveProducts(products []domain.Product) error {
	bytes, err := json.Marshal(products)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

// NewJsonStore crea un nuevo store de products
func NewStore(path string) Store {
	return &jsonStore{
		pathToFile: path,
	}
}

// GetAll devuelve todos los productos
func (s *jsonStore) GetAll() ([]domain.Product, error) {
	products, err := s.loadProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetById devuelve un producto por su id
func (s *jsonStore) GetByID(id int) (domain.Product, error) {
	products, err := s.loadProducts()
	if err != nil {
		return domain.Product{}, err
	}
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")
}

// Create agrega un nuevo producto
func (s *jsonStore) Create(product domain.Product) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	product.Id = len(products) + 1
	products = append(products, product)
	return s.saveProducts(products)
}

// Update actualiza un producto
func (s *jsonStore) Update(product domain.Product) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.Id == product.Id {
			products[i] = product
			return s.saveProducts(products)
		}
	}
	return errors.New("product not found")
}

// DeleteOne elimina un producto
func (s *jsonStore) Delete(id int) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.Id == id {
			products = append(products[:i], products[i+1:]...)
			return s.saveProducts(products)
		}
	}
	return errors.New("product not found")
}

// SearchPriceGt busca productos por precio mayor o igual que el precio dado
func (s *jsonStore) SearchPriceGt(price float64) []domain.Product {
	var productsFound []domain.Product
	products, err := s.loadProducts()
	if err != nil {
		return []domain.Product{}
	}
	for _, product := range products {
		if product.Price > price {
			productsFound = append(productsFound, product)
		}
	}
	return productsFound
}

// GetByCodeValue devuelve un producto por su code_value
func (s *jsonStore) GetByCodeValue(code string) (domain.Product, error) {
	products, err := s.loadProducts()
	if err != nil {
		return domain.Product{}, err
	}
	for _, product := range products {
		if product.CodeValue == code {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")
}

// Setea la cantidad de producto según la compra
// TODO QUE PASA CON EL HAPPY PATH
func (s *jsonStore) Buy(code string, quantity int) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.CodeValue == code && p.Quantity >= quantity {
			products[i].Quantity -= quantity
			return s.saveProducts(products)
		}
	}
	return errors.New("No se puede ejecutar la compra")

}
