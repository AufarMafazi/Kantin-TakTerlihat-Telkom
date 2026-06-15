package controllers

import (
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid", "details": err.Error()})
		return
	}

	orderID, err := h.orderService.Create(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderIDFormatted := fmt.Sprintf("ORD-%03d", orderID)

	c.JSON(http.StatusCreated, gin.H{
		"message":        "Pesanan sukses!",
		"order_id":       orderIDFormatted,
		"payment_status": order.PaymentStatus,
	})
}

func (h *OrderHandler) UpdatePaymentStatus(c *gin.Context) {
	orderID := c.Param("id")
	var input struct {
		PaymentStatus string `json:"payment_status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON salah"})
		return
	}

	if err := h.orderService.UpdatePayment(orderID, input.PaymentStatus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status pembayaran berhasil diperbarui"})
}

func (h *OrderHandler) GetUnpaidOrders(c *gin.Context) {
	orders, err := h.orderService.GetUnpaid()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func (h *OrderHandler) GetProcessingOrders(c *gin.Context) {
	orders, err := h.orderService.GetProcessing()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	orderID := c.Param("id")
	var input struct {
		OrderStatus string `json:"order_status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON salah"})
		return
	}

	if err := h.orderService.UpdateStatus(orderID, input.OrderStatus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status berhasil diperbarui"})
}