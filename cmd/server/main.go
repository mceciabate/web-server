package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mceciabate/web-server/cmd/server/employeeHandler"
	"github.com/mceciabate/web-server/cmd/server/productHandler"
	"github.com/mceciabate/web-server/internal/domain"
	"github.com/mceciabate/web-server/internal/employee"
	"github.com/mceciabate/web-server/internal/product"
)

func main() {
	var productsList = []domain.Product{}
	loadProducts("../data/products.json", &productsList)
	var employeesList = []domain.Employee{}
	loadEmployees("../data/employees.csv", &employeesList)
	//Instancio el repo y el service para productos
	repoP := product.NewRepository(productsList)
	serviceP := product.NewService(repoP)
	productHandler := productHandler.NewProductHandler(serviceP)

	//Instancio el repo y el service para employees
	repoE := employee.NewRepository(employeesList)
	serviceE := employee.NewService(repoE)
	employeeHandler := employeeHandler.NewEmployeeHandler(serviceE)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	r.GET("", func(c *gin.Context) { c.String(200, "Bienvenido a la empresa Gophers") })

	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAll())
		products.GET(":id", productHandler.GetByID())
		products.GET("/search", productHandler.Search())
		products.POST("", productHandler.Post())
		products.PUT(":id", productHandler.Put())
		products.DELETE(":id", productHandler.Delete())
		products.PATCH(":id", productHandler.Patch())
	}
	employees := r.Group("/employees")
	{
		employees.GET("", employeeHandler.GetAll())
		employees.GET(":id", employeeHandler.GetByID())
		employees.POST("", employeeHandler.Post())
		employees.PUT(":id", employeeHandler.Put())
		employees.DELETE(":id", employeeHandler.Delete())
	}

	r.Run(":8080")
}

// loadProducts carga los productos desde un archivo json
func loadProducts(path string, list *[]domain.Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}

// loadEmployess carga los employees desde un archivo csv
func loadEmployees(path string, list *[]domain.Employee) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}
