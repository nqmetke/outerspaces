package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	conn := DbInit()

	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("GET requet")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	r.POST("/posts/create", func(c *gin.Context) {
		var res NewPostJson

		c.BindJSON(&res)
		NewPost(conn, res.Title, res.Content)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")}
}
