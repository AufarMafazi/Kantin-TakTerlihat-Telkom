package services

import (
	"errors"
	"time"

	"backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("kantin_secret_key")

// Interface sebagai kontrak fungsi Auth
type AuthService interface {
	RegisterUser(user *models.User) error
	LoginUser(username, password string) (string, error)
}

type authService struct {
	db *gorm.DB
}

// Constructor untuk inisialisasi AuthService
func NewAuthService(db *gorm.DB) AuthService {
	return &authService{db: db}
}

// Logika Bisnis untuk Register
func (s *authService) RegisterUser(user *models.User) error {
	// Hashing password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("gagal memproses password")
	}
	user.Password = string(hashed)

	// Simpan ke DB
	if err := s.db.Create(user).Error; err != nil {
		return errors.New("gagal mendaftarkan user ke database")
	}

	return nil
}

// Logika Bisnis untuk Login
func (s *authService) LoginUser(username, password string) (string, error) {
	var user models.User

	// Cari user berdasarkan username
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("user tidak ditemukan")
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("password salah")
	}

	// Generasi JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("gagal membuat token keamanan")
	}

	return tokenString, nil
}