package models

import "time"

type Order struct {
    ID            uint        `json:"id" gorm:"primaryKey"`
    PaymentMethod string      `json:"payment_method" gorm:"column:payment_method"`
    PaymentStatus string      `json:"payment_status" gorm:"column:payment_status"`
    OrderStatus   string      `json:"order_status" gorm:"column:order_status"`
    TotalAmount   float64     `json:"total_amount" gorm:"column:total_amount"`
    Items         []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
    CreatedAt     time.Time   `json:"created_at"`
    UpdatedAt     time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderID   uint      `json:"order_id" gorm:"column:order_id"`
	MenuID    uint      `json:"menu_id" gorm:"column:menu_id"` 
	Menu      Menu      `json:"menu" gorm:"foreignKey:MenuID;references:ID"`
	Quantity  int       `json:"quantity" gorm:"column:quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}