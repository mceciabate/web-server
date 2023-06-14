package employee

import (
	"errors"

	"github.com/mceciabate/web-server/internal/domain"
)

type RepositoryE interface {
	GetAll() []domain.Employee
	GetByID(id int) (domain.Employee, error)
	Create(p domain.Employee) (domain.Employee, error)
	Update(p domain.Employee) error
}

type repositoryE struct {
	listEmployee []domain.Employee
}

// NewRepository crea un nuevo repositorio
func NewRepository(list []domain.Employee) RepositoryE {
	return &repositoryE{list}
}

// GetAll devuelve todos los empleados
func (r *repositoryE) GetAll() []domain.Employee {
	return r.listEmployee
}

// GetByID busca un empleado por su id
func (r *repositoryE) GetByID(id int) (domain.Employee, error) {
	for _, e := range r.listEmployee {
		if e.Id == id {
			return e, nil
		}
	}
	return domain.Employee{}, errors.New("employee not found")

}

// Create agrega un nuevo empleado
func (r *repositoryE) Create(e domain.Employee) (domain.Employee, error) {
	e.Id = len(r.listEmployee) + 1
	r.listEmployee = append(r.listEmployee, e)
	return e, nil
}

// Actualizar un empleado
func (r *repositoryE) Update(e domain.Employee) error {
	for i, e := range r.listEmployee {
		if e.Id == e.Id {
			r.listEmployee[i] = e
			return nil
		}
	}
	return errors.New("employee not found")
}
