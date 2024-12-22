package handler

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/drunkleen/go-url-shortner/config"
	"github.com/drunkleen/go-url-shortner/shortener"
	"github.com/drunkleen/go-url-shortner/store"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UrlCreationRequest struct {
	LongUrl string `json:"url" binding:"required"`
	UserId  string `json:"user_id" binding:""`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	creationRequest.UserId = generateUUIDFromIP(getClientIP(c))
	log.Printf("User ID: %s", creationRequest.UserId)
	log.Printf("---- Provided IP: %s ----", getClientIP(c))

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": fmt.Sprintf("http://localhost%s/", config.Port) + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}

func getClientIP(c *gin.Context) string {
	xForwardedFor := c.GetHeader("X-Forwarded-For")
	if xForwardedFor != "" {
		ips := strings.Split(xForwardedFor, ",")
		return strings.TrimSpace(ips[0])
	}

	xRealIP := c.GetHeader("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}

	ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		return c.Request.RemoteAddr
	}
	return ip
}

func generateUUIDFromIP(ip string) string {
	namespace := uuid.NewMD5(uuid.NameSpaceDNS, []byte("www.drunkleen.com"))
	return uuid.NewMD5(namespace, []byte(ip)).String()
}
