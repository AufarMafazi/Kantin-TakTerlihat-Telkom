package main

import (
	"log"
	"order-service/handlers"
	"order-service/models"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Koneksi Database
	dsn := "root:@tcp(127.0.0.1:3306)/tugasbesariae?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	db.AutoMigrate(&models.Order{}, &models.OrderItem{})

	// Koneksi NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("Gagal konek ke NATS:", err)
	}
	defer nc.Close()

	// Setup Handler 
	orderHandler := &handlers.OrderHandler{
		DB:       db,
		NatsConn: nc,
	}

	r := gin.Default()
	r.POST("/orders", orderHandler.CreateOrder)
	r.PATCH("/orders/:id/status", func(c *gin.Context) { 
		handlers.UpdatePaymentStatus(c, db, nc) 
	})
	r.GET("/orders/pending", func(c *gin.Context) { 
    	handlers.GetUnpaidOrders(c, db) 
	})

	r.GET("/orders/kitchen", func(c *gin.Context) { 
		handlers.GetProcessingOrders(c, db) 
	})

	r.PATCH("/orders/:id/done", func(c *gin.Context) { 
		handlers.UpdateOrderStatus(c, db, nc) 
	})


	log.Println("Order Service berjalan di port 8081...")
	r.Run(":8081")
}