package handler

import (
	"net/http"
	"back-end/internal/hub"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket" // Pastikan import ini ada
)

// Tambahkan variabel upgrader ini di bagian atas file
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Mengizinkan koneksi dari mana saja (frontend Vue)
	},
}

type OrderHandler struct {
	Hub *hub.Hub
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
    // ... isi kode CreateOrder Anda ...
}

// TAMBAHKAN FUNGSI INI:
func (h *OrderHandler) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	
	// Masukkan koneksi kasir ke daftar klien di Hub
	h.Hub.Mutex.Lock()
	h.Hub.Clients[conn] = true
	h.Hub.Mutex.Unlock()
}