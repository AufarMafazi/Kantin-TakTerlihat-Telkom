package main

import (
	"backend/config"
	"backend/internal/controllers"
	"backend/internal/services"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDatabase()

	notifService := services.NewNotificationService()
	notifHandler := controllers.NewNotificationHandler(notifService)

	authService := services.NewAuthService(db)
	authHandler := controllers.NewAuthHandler(authService)

	menuService := services.NewMenuService(db)
	menuHandler := controllers.NewMenuHandler(menuService)

	orderService := services.NewOrderService(db, notifService)
	orderHandler := controllers.NewOrderHandler(orderService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/ws", notifHandler.HandleWS)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	r.GET("/menu", menuHandler.GetMenu)

	r.POST("/orders", orderHandler.CreateOrder)
	r.PATCH("/orders/:id/status", orderHandler.UpdatePaymentStatus)
	r.GET("/orders/pending", orderHandler.GetUnpaidOrders)
	r.GET("/orders/kitchen", orderHandler.GetProcessingOrders)
	r.PATCH("/orders/:id/done", orderHandler.UpdateOrderStatus)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}