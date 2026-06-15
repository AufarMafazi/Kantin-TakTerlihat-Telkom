package services

import (
	"log"
	"sync"
	"github.com/gorilla/websocket"
)

type NotificationService interface {
	RegisterClient(conn *websocket.Conn)
	UnregisterClient(conn *websocket.Conn)
	BroadcastMessage(messageType int, data []byte)
}

type notificationService struct {
	clients map[*websocket.Conn]bool
	mutex   sync.Mutex
}

func NewNotificationService() NotificationService {
	return &notificationService{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (s *notificationService) RegisterClient(conn *websocket.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.clients[conn] = true
	log.Printf("Client WebSocket baru terhubung: %s", conn.RemoteAddr())
}

func (s *notificationService) UnregisterClient(conn *websocket.Conn) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, exists := s.clients[conn]; exists {
		delete(s.clients, conn)
		conn.Close()
		log.Printf("Client WebSocket terputus: %s", conn.RemoteAddr())
	}
}

func (s *notificationService) BroadcastMessage(messageType int, data []byte) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Printf("Broadcasting via WS: %s", string(data))
	for client := range s.clients {
		err := client.WriteMessage(messageType, data)
		if err != nil {
			log.Printf("Gagal mengirim pesan ke %s, menutup koneksi: %v", client.RemoteAddr(), err)
			go s.UnregisterClient(client)
		}
	}
}