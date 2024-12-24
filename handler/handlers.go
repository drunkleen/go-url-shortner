package handler

import (
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/drunkleen/go-url-shortner/config"
	"github.com/drunkleen/go-url-shortner/shortener"
	"github.com/drunkleen/go-url-shortner/store"
	"github.com/drunkleen/go-url-shortner/utils"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"url" binding:"required"`
	UserId  string
}

// CreateShortUrl is a Gin handler function that creates a short URL given a long URL and saves it into the store.
// It returns the created short URL as a JSON response.
func CreateShortUrl(c *gin.Context) {
	// The request body is expected to contain the long URL as a JSON object.
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		// If the request body is invalid, return a Bad Request response with the error message.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a UUID from the client IP address.
	creationRequest.UserId = utils.GenerateUUIDFromIP(getClientIP(c))

	// Generate a short URL given the long URL and the generated UUID.
	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)

	// Save the short URL mapping into the store.
	if err := store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId); err != nil {
		// If an error occurs while saving the mapping, log the error and return an Internal Server Error response.
		log.Printf("Failed to save url mapping: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save url mapping"})
		return
	}

	// Return the created short URL as a JSON response.
	c.JSON(http.StatusCreated, gin.H{
		"message":   "short url created successfully",
		"short_url": config.AppConfig.Host + config.AppConfig.Port + "/" + shortUrl,
	})
}

// HandleShortUrlRedirect is a Gin handler function that redirects the user to the original URL using the short URL as a parameter.
// It retrieves the original URL from the store using the provided short URL, and then redirects the user to that URL.
func HandleShortUrlRedirect(c *gin.Context) {
	// Extract the short URL from the request parameters.
	shortUrl := c.Param("shortUrl")

	// Retrieve the initial/original URL from the store using the short URL.
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	if initialUrl == "" {
		// If the mapping could not be found, return a Not Found response.
		c.JSON(http.StatusNotFound, gin.H{"error": "Url not found"})
		return
	}

	if !strings.HasPrefix(initialUrl, "http://") && !strings.HasPrefix(initialUrl, "https://") {
		initialUrl = "https://" + initialUrl
	}

	// Trim any whitespace from the initial URL.
	initialUrl = strings.TrimSpace(initialUrl)

	// Redirect the user to the original URL with a 302 status code.
	c.Redirect(http.StatusPermanentRedirect, initialUrl)
}

// getClientIP retrieves the client's IP address from the request headers or remote address.
// It prioritizes the "X-Forwarded-For" header, followed by the "X-Real-IP" header, and finally the remote address.
func getClientIP(c *gin.Context) string {
	// Check the "X-Forwarded-For" header for the client's IP address.
	xForwardedFor := c.GetHeader("X-Forwarded-For")
	if xForwardedFor != "" {
		// Split the header value by commas and return the first IP address.
		ips := strings.Split(xForwardedFor, ",")
		return strings.TrimSpace(ips[0])
	}

	// Check the "X-Real-IP" header for the client's IP address.
	xRealIP := c.GetHeader("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}

	// If no headers are present, use the remote address from the request.
	ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		// Return the raw remote address if there's an error parsing it.
		return c.Request.RemoteAddr
	}
	return ip
}
