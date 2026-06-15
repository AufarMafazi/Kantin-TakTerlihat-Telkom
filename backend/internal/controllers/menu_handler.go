package controllers

import (
	"backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	menuService services.MenuService
}

func NewMenuHandler(menuService services.MenuService) *MenuHandler {
	return &MenuHandler{menuService: menuService}
}

func (h *MenuHandler) GetMenu(c *gin.Context) {
	menus, err := h.menuService.GetAllMenus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menus)
}