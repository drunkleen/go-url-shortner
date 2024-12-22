package main

import (
	"log"

	"github.com/drunkleen/go-url-shortner/config"
	"github.com/drunkleen/go-url-shortner/handler"
	"github.com/drunkleen/go-url-shortner/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStoreService()

	log.Printf("Server is running on port %s", config.Port)
	if err := r.Run(config.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
