package config

import (
	"fmt"
	"log"

	"backend/internal/models" 

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	user := "root"
	password := ""
	host := "127.0.0.1"
	port := "3306"
	dbName := "tugasbesariae"   

	// 2. Susun DSN dengan format khusus MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)

	// 3. Buka koneksi ke database menggunakan GORM dengan driver MySQL
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database MySQL %s: %v", dbName, err)
	}

	fmt.Printf("Koneksi database MySQL '%s' berhasil!\n", dbName)

	// 4. Auto Migration (SUDAH DIUBAH)
	// Kita masukkan semua struct model ke sini agar GORM otomatis membuat/memperbarui tabelnya
	err = database.AutoMigrate(
		&models.User{}, 
		&models.Order{}, 
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi database: %v", err)
	}

	fmt.Println("Migrasi seluruh tabel ke MySQL berhasil!")

	// Simpan ke variabel global dan kembalikan nilainya
	DB = database
	return database
}