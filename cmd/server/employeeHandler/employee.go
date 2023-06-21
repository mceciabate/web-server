package employeeHandler

import (
	"errors"
	"net/http"
	"os"
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

// GetByID obtiene un empleado por su id
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

// Delete elimina un empleado
func (h *employeeHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"message": "product deleted"})
	}
}

// Patch update selected fields of a product WIP
func (h *employeeHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name   string `json:"name,omitempty"`
		Active bool   `json:"is_active,omitempty"`
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
		var r Request
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(400, gin.H{"error": "invalid request"})
			return
		}
		update := domain.Employee{
			Name:   r.Name,
			Active: r.Active,
		}
		if update.Name != "" {
			valid, err := validateEmptys(&update)
			if !valid {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
		}
		e, err := h.s.Update(id, update)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, e)
	}
}

//Filtrar empleados activos

func (h *employeeHandler) GetActives() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employees, _ := h.s.FilterActive()
		ctx.JSON(200, employees)

	}
}
