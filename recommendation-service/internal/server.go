package internal

import (
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	// Health check endpoint
	r.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
