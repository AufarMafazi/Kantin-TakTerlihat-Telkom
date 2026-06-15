package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
	"sync"
)

var (
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	clients  = make(map[*websocket.Conn]bool)
	mutex    sync.Mutex
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)

	nc.Subscribe("order.created", func(m *nats.Msg) {
		log.Printf("Notifikasi diterima: %s", string(m.Data))
		
		// Broadcast ke semua device yang terhubung
		mutex.Lock()
		for client := range clients {
			client.WriteMessage(websocket.TextMessage, m.Data)
		}
		mutex.Unlock()
	})

	nc.Subscribe("order.status_updated", func(m *nats.Msg) {
		log.Printf("Event diterima dari NATS: %s", string(m.Data))
		mutex.Lock()
		for client := range clients {
			// Kirim sinyal sederhana agar frontend melakukan fetch ulang
			client.WriteMessage(websocket.TextMessage, []byte("REFRESH"))
		}
		mutex.Unlock()
	})

	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
        conn, _ := upgrader.Upgrade(c.Writer, c.Request, nil)
        
        mutex.Lock()
        clients[conn] = true
        mutex.Unlock()

        // Jalankan goroutine untuk menunggu pesan (menandakan koneksi hidup)
        // Saat client putus, goroutine ini akan berhenti
        go func() {
            defer func() {
                mutex.Lock()
                delete(clients, conn)
                mutex.Unlock()
                conn.Close()
            }()
            
            // Loop ini menjaga koneksi tetap hidup
            for {
                if _, _, err := conn.ReadMessage(); err != nil {
                    break // Keluar dari loop saat client disconnect
                }
            }
        }()
    })

	r.Run(":8082")
}