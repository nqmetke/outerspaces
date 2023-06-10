package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()

	GormInit()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"connected": "true",
		})
	})

	r.POST("/api/register", func(c *gin.Context) {
		UserRegister(c)
	})

	r.POST("/api/login", func(c *gin.Context) {
		UserLogin(c)
	})

	protected := r.Group("/api/admin")
	protected.Use(JwtAuthMiddleware())
	protected.GET("/user", GetCurrentUser)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")}
}
