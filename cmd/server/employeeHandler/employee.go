package employeeHandler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mceciabate/web-server/internal/domain"
	"github.com/mceciabate/web-server/internal/employee"
)

type employeeHandler struct {
	s employee.ServiceE
}

// NewProductHandler crea un nuevo controller de productos
func NewEmployeeHandler(s employee.ServiceE) *employeeHandler {
	return &employeeHandler{
		s: s,
	}
}

// GetAll obtiene todos los empleados
func (h *employeeHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, _ := h.s.GetAll()
		c.JSON(200, products)
	}
}

// GetByID obtiene un producto por su id
func (h *employeeHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		employee, err := h.s.GetByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "employee not found"})
			return
		}
		c.JSON(200, employee)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptys(employee *domain.Employee) (bool, error) {
	if employee.Name == "" {
		return false, errors.New("name field can't be empty")
	}
	return true, nil
}

// Post crear un producto nuevo
func (h *employeeHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var employee domain.Employee
		err := ctx.ShouldBindJSON(&employee)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid employee"})
			return
		}
		valid, err := validateEmptys(&employee)
		if !valid {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		e, err := h.s.Create(employee)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, e)
	}
}

func (h *employeeHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid id",
			})
			return
		}
		var employee domain.Employee
		err = c.ShouldBindJSON(&employee)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid body",
			})
			return
		}
		valid, err := validateEmptys(&employee)
		if !valid {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		e, err := h.s.Update(id, employee)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err,
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, e)
	}
}
