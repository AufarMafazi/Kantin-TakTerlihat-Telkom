package main

import (
	"database/sql"
	"log"
	"menu-service/internal/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Koneksi ke Database
	// Ganti "user:password@tcp(localhost:3306)/kantin_db" dengan kredensial Anda
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/tugasbesariae")
	if err != nil {
		log.Fatal("Koneksi DB gagal: ", err)
	}
	defer db.Close()

	// Inisialisasi Handler
	menuHandler := &handler.MenuHandler{DB: db}

	// Setup Router
	r := gin.New() // Gunakan New() agar kita bisa konfigurasi manual
    r.Use(gin.Recovery()) // Menambahkan middleware Recovery standar
    r.Use(gin.Logger())   // Menambahkan middleware Logger standar
    
    // MATIKAN FITUR REDIRECT OTOMATIS:
    r.RedirectTrailingSlash = false 

    // Inisialisasi Handler & Rute
    r.GET("/menu", menuHandler.GetMenu) // Pastikan rutenya sesuai
    r.GET("/menu/", menuHandler.GetMenu) // Opsional: tambahkan ini jika ingin support kedua format

	// Jalankan Service
	log.Println("Menu Service berjalan di port 8083...")
	r.Run(":8083")
}