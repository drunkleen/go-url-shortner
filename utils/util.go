package utils

import (
	"github.com/drunkleen/go-url-shortner/config"
	"github.com/google/uuid"
)

// generateUUIDFromIP generates a UUID from the client's IP address.
// It uses the MD5 version 5 UUID algorithm to generate a UUID that is
// unique for each client IP address.
func GenerateUUIDFromIP(ip string) string {
	// Create a namespace UUID from the domain name defined in the application configuration.
	namespace := uuid.NewMD5(uuid.NameSpaceDNS, []byte(config.AppConfig.Host))

	// Generate a UUID using the MD5 algorithm and the client IP address.
	// This UUID is intended to uniquely identify the client based on their IP address.
	return uuid.NewMD5(namespace, []byte(ip)).String()
}
