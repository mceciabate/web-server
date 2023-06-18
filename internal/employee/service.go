package employee

import "github.com/mceciabate/web-server/internal/domain"

type ServiceE interface {
	GetAll() ([]domain.Employee, error)
	GetByID(id int) (domain.Employee, error)
	Create(e domain.Employee) (domain.Employee, error)
	Update(id int, e domain.Employee) (domain.Employee, error)
	Delete(id int) error
}

type serviceE struct {
	r RepositoryE
}

// NewService crea un nuevo servicio
func NewService(r RepositoryE) ServiceE {
	return &serviceE{r}
}

// GetAll devuelve todos los empleado
func (s *serviceE) GetAll() ([]domain.Employee, error) {
	lE := s.r.GetAll()
	return lE, nil
}

// GetByID busca un producto por su id
func (s *serviceE) GetByID(id int) (domain.Employee, error) {
	e, err := s.r.GetByID(id)
	if err != nil {
		return domain.Employee{}, err
	}
	return e, nil
}

// Create agrega un nuevo producto
func (s *serviceE) Create(e domain.Employee) (domain.Employee, error) {
	e, err := s.r.Create(e)
	if err != nil {
		return domain.Employee{}, err
	}
	return e, nil
}

func (s *serviceE) Update(id int, e domain.Employee) (domain.Employee, error) {
	_, err := s.r.GetByID(id)
	if err != nil {
		return domain.Employee{}, err
	}

	e.Id = id
	err = s.r.Update(e)
	if err != nil {
		return domain.Employee{}, err
	}
	return e, nil
}

// Delete elimina un producto
func (s *serviceE) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
