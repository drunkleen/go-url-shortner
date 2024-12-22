package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

const (
	PORT string = ":8080"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	log.Printf("Server is running on port %s", PORT)
	if err := r.Run(PORT); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
