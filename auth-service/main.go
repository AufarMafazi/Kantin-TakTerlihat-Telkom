package main

import (
	"auth-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db := InitDB()
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) { handlers.Register(c, db) })
	r.POST("/login", func(c *gin.Context) { handlers.Login(c, db) })

	r.Run(":8084")
}