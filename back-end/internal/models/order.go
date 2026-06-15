package models

// OrderItem mewakili setiap item di dalam cart
type OrderItem struct {
	ItemID   int `json:"item_id"`
	Quantity int  `json:"quantity"`
}

// OrderRequest adalah struct yang menangkap body JSON dari frontend
type OrderRequest struct {
	Items         []OrderItem `json:"items"`
	TotalAmount   float64     `json:"total_amount"`
	PaymentMethod string      `json:"payment_method"`
}	