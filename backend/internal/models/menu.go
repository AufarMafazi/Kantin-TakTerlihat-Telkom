package models

type Menu struct {
	ID    uint    `json:"id" gorm:"primaryKey;type:bigint unsigned"` 
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Pic   string  `json:"pic"`
}

func (Menu) TableName() string {
	return "menu"
}