package main

import (
	"net/http"
	"back-end/internal/database"
	"back-end/internal/handler"
	"back-end/internal/hub"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

// Fungsi untuk membungkus setiap request agar mengizinkan CORS
// func enableCORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Mengizinkan frontend (Vue) mengakses server ini
// 		w.Header().Set("Access-Control-Allow-Origin", "*") 
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 		// Jika request-nya adalah OPTIONS (pre-flight), langsung return
// 		if r.Method == "OPTIONS" {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

func main() {
db := database.InitDB()
    defer db.Close()
    r := gin.Default()
    r.Use(cors.Default()) // Pastikan CORS ada di sini

    // 2. Hub/Handler
    orderHub := hub.NewHub()
    go orderHub.Run()
    orderHandler := &handler.OrderHandler{Hub: orderHub}

    // 3. DAFTARKAN SEMUA RUTE KE 'r' (GIN)
    r.POST("/orders", orderHandler.CreateOrder)
    r.GET("/ws", orderHandler.HandleWebSocket)
    r.GET("/menu", gin.WrapH(http.HandlerFunc(handler.HandleGetMenu(db))))

    // 4. JALANKAN (Hanya ini, jangan ada ListenAndServe lain!)
    r.Run(":8080")
}