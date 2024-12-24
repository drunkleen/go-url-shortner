package main

import (
	"log"

	"github.com/drunkleen/go-url-shortner/config"
	"github.com/drunkleen/go-url-shortner/handler"
	"github.com/drunkleen/go-url-shortner/store"
	"github.com/gin-gonic/gin"
)

// main initializes the URL Shortener API server. It loads the configuration,
// sets up the Gin router with defined routes for creating and redirecting short URLs,
// initializes the store service, and starts the server on the configured port.
func main() {
	// Load the application configuration
	config.LoadConfig()

	// Initialize the Gin router
	r := gin.Default()

	// Define a GET route to serve a welcome message
	// This route is used to test the API.
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	// Define a POST route to create a short URL
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	// Define a GET route to handle short URL redirection
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Initialize the store service for URL mapping
	store.InitializeStoreService()

	// Start the server and listen on the configured port
	log.Printf(">> Server is running on port %s", config.AppConfig.Port)
	if err := r.Run(config.AppConfig.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
