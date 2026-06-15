package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"order-service/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type OrderHandler struct {
	DB       *gorm.DB
	NatsConn *nats.Conn
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Data tidak valid",
			"details": err.Error(),
		})
		return
	}

	if strings.ToUpper(order.PaymentMethod) == "QRIS" {
		order.PaymentStatus = "Lunas"
	} else {
		order.PaymentStatus = "Menunggu Pembayaran"
	}

	// 2. Simpan ke Database
	if err := h.DB.Create(&order).Error; err != nil {
		log.Printf("Gagal simpan ke database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal simpan database"})
		return
	}

	// 3. Format ID untuk Notifikasi (Contoh: ORD-001)
	orderIDFormatted := fmt.Sprintf("ORD-%03d", order.ID)

	// Encode data
	data, _ := json.Marshal(map[string]interface{}{
		"formatted_id":   orderIDFormatted,
		"items":          order.Items,
		"total":          order.TotalAmount,
		"payment_status": order.PaymentStatus, 
	})

	// Publish ke NATS dengan Error Handling
	if err := h.NatsConn.Publish("order.created", data); err != nil {
		log.Printf("Gagal publish ke NATS: %v", err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":        "Pesanan sukses!",
		"order_id":       orderIDFormatted,
		"payment_status": order.PaymentStatus,
	})
}

func UpdatePaymentStatus(c *gin.Context, db *gorm.DB, nc *nats.Conn) {
	orderID := c.Param("id")
	
	var input struct {
		PaymentStatus string `json:"payment_status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON salah"})
		return
	}

	if err := db.Model(&models.Order{}).Where("id = ?", orderID).Update("payment_status", input.PaymentStatus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update status pembayaran"})
		return
	}

	// 2. Publish ke NATS agar sistem tahu ada perubahan (Broadcast)
	// Kita gunakan topik yang sama dengan status dapur agar frontend cukup me-refresh data
	payload := fmt.Sprintf(`{"id": %s, "payment_status": "%s"}`, orderID, input.PaymentStatus)
	nc.Publish("order.status_updated", []byte(payload))

	c.JSON(http.StatusOK, gin.H{"message": "Status pembayaran berhasil diperbarui"})
}

func GetUnpaidOrders(c *gin.Context, db *gorm.DB) {
	var orders []models.Order
	err := db.Where("payment_status = ?", "Menunggu Pembayaran").
		Preload("Items"). 
		Find(&orders).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data pesanan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func GetProcessingOrders(c *gin.Context, db *gorm.DB) {
	var orders []models.Order

	err := db.Where("payment_status = ? AND order_status = ?", "Lunas", "Sedang Diproses").
		Preload("Items"). 
		Find(&orders).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data pesanan dapur"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func UpdateOrderStatus(c *gin.Context, db *gorm.DB, nc *nats.Conn) {
	orderID := c.Param("id")

	// Struct input
	var input struct {
		OrderStatus string `json:"order_status" binding:"required"`
	}

	// Bind JSON dari request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON salah: " + err.Error()})
		return
	}

	// Update Database
	result := db.Model(&models.Order{}).Where("id = ?", orderID).Update("order_status", input.OrderStatus)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update status database"})
		return
	}

	// Publish ke NATS untuk trigger WebSocket
	payload := fmt.Sprintf(`{"id": "%s", "order_status": "%s"}`, orderID, input.OrderStatus)
	err := nc.Publish("order.status_updated", []byte(payload))
	if err != nil {
		fmt.Println("Gagal publish ke NATS:", err)
	}

	// Response Sukses
	c.JSON(http.StatusOK, gin.H{"message": "Status berhasil diperbarui"})
}

