package store

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/drunkleen/go-url-shortner/config"
	"github.com/redis/go-redis/v9"
)

var (
	storeService  = &StoreService{}
	ctx           = context.Background()
	CacheDuration time.Duration
)

// StoreService provides methods to interact with the Redis store.
type StoreService struct {
	redisClient *redis.Client
}

// InitializeStoreService initializes the StoreService singleton with the Redis client.
// It sets the CacheDuration variable based on the configured value, and creates a new Redis client.
func InitializeStoreService() *StoreService {
	var cacheDurationInt int
	if config.AppConfig.CacheDuration == "" {
		cacheDurationInt = 1440
		config.AppConfig.CacheDuration = "1440"
		log.Printf("CacheDuration not set, using default value: %d", cacheDurationInt)
	} else {
		var err error
		cacheDurationInt, err = strconv.Atoi(config.AppConfig.CacheDuration)
		if err != nil {
			log.Fatalf("Error parsing CacheDuration: %v", err)
		}
	}

	// Set the CacheDuration variable to the parsed value.
	CacheDuration = time.Duration(cacheDurationInt) * time.Minute

	// Create a new Redis client.
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisURL + ":" + config.AppConfig.RedisPort,
		Password: config.AppConfig.RedisPassword,
		DB:       0,
	})

	// Ping the Redis server to check the connection.
	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		log.Panicf("Failed to connect to Redis: %v", err)
	}

	log.Printf(">> Redis connected successfully at %s:%s", config.AppConfig.RedisURL, config.AppConfig.RedisPort)
	// Set the Redis client to the StoreService singleton.
	storeService.redisClient = redisClient
	return storeService
}

// SaveUrlMapping stores the mapping between a short URL and its original long URL in the Redis store.
// The mapping is associated with a specific user ID and is set to expire after the configured cache duration.
//
// Parameters:
//
//	shortUrl - the short URL to be stored
//	longUrl - the original URL associated with the short URL
//	userId - the user ID associated with the short URL
//
// Returns an error if the mapping could not be stored.
func SaveUrlMapping(shortUrl, longUrl, userId string) error {
	// Attempt to set the short URL and long URL mapping in the Redis store with an expiration duration.
	if err := storeService.redisClient.Set(ctx, shortUrl, longUrl, CacheDuration).Err(); err != nil {
		// If an error occurs, return the error.
		return err
	}
	// If the mapping was stored successfully, return nil.
	return nil
}

// RetrieveInitialUrl retrieves the original URL from the Redis store given a short URL.
// It returns an empty string if the mapping could not be found.
func RetrieveInitialUrl(shortUrl string) string {
	// Attempt to get the original URL associated with the given short URL from the Redis store.
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		// If the mapping could not be found, log an error message and return an empty string.
		log.Printf("Failed to retrieve initial url: %v", err)
		return ""
	}
	return result
}
