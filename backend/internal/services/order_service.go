package services

import (
	"backend/internal/models"
	"errors"
	"log"
	"strings"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type OrderService interface {
	Create(order *models.Order) (uint, error)
	UpdatePayment(orderID string, status string) error
	GetUnpaid() ([]models.Order, error)
	GetProcessing() ([]models.Order, error)
	UpdateStatus(orderID string, status string) error
}

type orderService struct {
	db       *gorm.DB
	notifSvc NotificationService
}

func NewOrderService(db *gorm.DB, notifSvc NotificationService) OrderService {
	return &orderService{
		db:       db,
		notifSvc: notifSvc,
	}
}

func (s *orderService) Create(order *models.Order) (uint, error) {
	if strings.ToUpper(order.PaymentMethod) == "QRIS" {
		order.PaymentStatus = "Lunas"
	} else {
		order.PaymentStatus = "Menunggu Pembayaran"
	}

	order.OrderStatus = "Sedang Diproses"

	if err := s.db.Create(order).Error; err != nil {
		log.Printf("Gagal menyimpan order ke database: %v", err)
		return 0, errors.New("gagal menyimpan pesanan")
	}

	s.notifSvc.BroadcastMessage(websocket.TextMessage, []byte("PESANAN_BARU"))

	return order.ID, nil
}

func (s *orderService) UpdatePayment(orderID string, status string) error {
	var order models.Order
	
	if err := s.db.First(&order, orderID).Error; err != nil {
		return errors.New("pesanan tidak ditemukan")
	}

	if err := s.db.Model(&order).Update("payment_status", status).Error; err != nil {
		log.Printf("Gagal mengupdate status pembayaran order %s: %v", orderID, err)
		return errors.New("gagal update status pembayaran di database")
	}

	s.notifSvc.BroadcastMessage(websocket.TextMessage, []byte("REFRESH"))
	return nil
}

func (s *orderService) GetUnpaid() ([]models.Order, error) {
	var orders []models.Order
	
	err := s.db.Where("payment_status = ?", "Menunggu Pembayaran").Preload("Items.Menu").Find(&orders).Error
	if err != nil {
		log.Printf("Gagal mengambil data unpaid orders: %v", err)
		return nil, errors.New("gagal mengambil data pesanan belum dibayar")
	}
	
	return orders, nil
}

func (s *orderService) GetProcessing() ([]models.Order, error) {
	var orders []models.Order
	
	err := s.db.Where("payment_status = ? AND order_status = ?", "Lunas", "Sedang Diproses").Preload("Items.Menu").Find(&orders).Error
	if err != nil {
		log.Printf("Gagal mengambil data processing orders untuk dapur: %v", err)
		return nil, errors.New("gagal mengambil data pesanan dapur")
	}
	
	return orders, nil
}

func (s *orderService) UpdateStatus(orderID string, status string) error {
	var order models.Order

	if err := s.db.First(&order, orderID).Error; err != nil {
		return errors.New("pesanan tidak ditemukan")
	}

	if err := s.db.Model(&order).Update("order_status", status).Error; err != nil {
		log.Printf("Gagal mengupdate status order %s: %v", orderID, err)
		return errors.New("gagal update status order di database")
	}

	s.notifSvc.BroadcastMessage(websocket.TextMessage, []byte("REFRESH"))
	return nil
}