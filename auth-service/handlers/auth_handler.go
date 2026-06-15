package handlers

import (
	"auth-service/models"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("kantin_secret_key")

func Register(c *gin.Context, db *gorm.DB) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashed)
	db.Create(&user)
	c.JSON(201, gin.H{"message": "User terdaftar"})
}

func Login(c *gin.Context, db *gorm.DB) {
	var input models.User
	var user models.User
	c.ShouldBindJSON(&input)
	
	if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "User tidak ditemukan"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Password salah"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": user.Username, "exp": time.Now().Add(24 * time.Hour).Unix()})
	tokenString, _ := token.SignedString(jwtKey)
	c.JSON(200, gin.H{"token": tokenString})
}