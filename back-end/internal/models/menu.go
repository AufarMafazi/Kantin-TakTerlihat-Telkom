package models

type Menu struct {
    ID    string  `json:"id"`
    Name  string  `json:"name"`
    Price int 	`json:"price"`
    Pic string `json:"pic"`
}