package main

import (
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Fungsi pembantu untuk memproses request HTTP ke service lain
func serveReverseProxy(target string, c *gin.Context) {
	targetURL, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	c.Request.URL.Host = targetURL.Host
	c.Request.URL.Scheme = targetURL.Scheme
	
	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.RedirectTrailingSlash = false

	handleWS := func(c *gin.Context) {
		conn, err := net.Dial("tcp", "localhost:8082")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Notification service unavailable"})
			return
		}
		c.Request.Write(conn)

		hijacker, ok := c.Writer.(http.Hijacker)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		
		clientConn, _, err := hijacker.Hijack()
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		go io.Copy(conn, clientConn)
		go io.Copy(clientConn, conn)
	}

	// Routing
	// Service WS
	r.Any("/ws", handleWS)
	r.Any("/ws/*any", handleWS)

	// service menu
	r.Any("/menu", func(c *gin.Context) { serveReverseProxy("http://localhost:8083", c)})
	r.Any("/menu/*any", func(c *gin.Context) { serveReverseProxy("http://localhost:8083", c)})

	// service order
	r.Any("/orders", func(c *gin.Context) { serveReverseProxy("http://localhost:8081", c)})
	r.Any("/orders/*any", func(c *gin.Context) { serveReverseProxy("http://localhost:8081", c)})

	// service auth
	r.Any("/login", func(c *gin.Context) { serveReverseProxy("http://localhost:8084", c) })
	r.Any("/register", func(c *gin.Context) { serveReverseProxy("http://localhost:8084", c) })

	r.Run(":8080")
}