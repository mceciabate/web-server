package domain

type Purchase struct {
	CodeValue  string  `json:"code_value" binding:"required"`
	Quantity   int     `json:"quantity" binding:"required"`
	TotalPrice float64 `json:"total_price" binding:"required"`
}
