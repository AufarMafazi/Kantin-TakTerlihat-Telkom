package hub

import (
	"sync"
	"github.com/gorilla/websocket"
)

// Hub menjaga kumpulan koneksi klien aktif
type Hub struct {
	// Clients menyimpan daftar koneksi kasir yang sedang online
	Clients   map[*websocket.Conn]bool
	// Broadcast adalah saluran untuk mengirim pesan ke semua kasir
	Broadcast chan interface{}
	// Mutex untuk memastikan operasi pada map aman dari race condition
	Mutex     sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan interface{}),
	}
}

func (h *Hub) Run() {
	for {
		// Mengambil pesan dari channel Broadcast
		msg := <-h.Broadcast
		
		// Mengirim pesan ke semua kasir yang terhubung
		h.Mutex.Lock()
		for client := range h.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(h.Clients, client)
			}
		}
		h.Mutex.Unlock()
	}
}