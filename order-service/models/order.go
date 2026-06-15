package models

import "gorm.io/gorm"

type Order struct {
	ID            uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	Items         []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	TotalAmount   float64     `json:"total_amount"`
	PaymentMethod string      `json:"payment_method"`
	PaymentStatus string      `json:"payment_status"`
	OrderStatus   string        `json:"order_status" gorm:"default:'Sedang Diproses'"`
}

type OrderItem struct {
	gorm.Model
	OrderID            uint    `json:"-"`
	Name               string  `json:"name"`
	Quantity           int     `json:"quantity"`
	PriceAtTransaction float64 `json:"price" gorm:"column:price_at_transaction"`
}