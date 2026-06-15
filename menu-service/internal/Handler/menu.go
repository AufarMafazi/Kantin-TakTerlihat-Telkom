package handler

import (
	"database/sql"
	"menu-service/internal/repository"
	"net/http"
	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	DB *sql.DB
}

func (h *MenuHandler) GetMenu(c *gin.Context) {
    menus, err := repository.FetchMenus(h.DB)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, menus)
}