package domain

type Employee struct {
	Id     int    `json:"is"`
	Name   string `json:"name" binding:"required"`
	Active bool   `json:"is_active" binding:"required"`
}
