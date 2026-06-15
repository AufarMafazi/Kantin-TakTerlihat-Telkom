package controllers

import (
	"backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type NotificationHandler struct {
	notifSvc services.NotificationService
	upgrader websocket.Upgrader
}

func NewNotificationHandler(notifSvc services.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		notifSvc: notifSvc,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
	}
}

func (h *NotificationHandler) HandleWS(c *gin.Context) {
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal inisialisasi koneksi WebSocket"})
		return
	}
	h.notifSvc.RegisterClient(conn)

	go func() {
		defer h.notifSvc.UnregisterClient(conn)
		
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				break
			}
		}
	}()
}