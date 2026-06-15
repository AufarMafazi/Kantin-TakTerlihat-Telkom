package services

import (
	"backend/internal/models"
	"gorm.io/gorm"
)

type MenuService interface {
	GetAllMenus() ([]models.Menu, error)
}

type menuService struct {
	db *gorm.DB
}

func NewMenuService(db *gorm.DB) MenuService {
	return &menuService{db: db}
}

func (s *menuService) GetAllMenus() ([]models.Menu, error) {
    var menus []models.Menu
    
    err := s.db.Raw("SELECT id, name, price, pic FROM menu").Scan(&menus).Error
    if err != nil {
        return nil, err
    }

    return menus, nil
}