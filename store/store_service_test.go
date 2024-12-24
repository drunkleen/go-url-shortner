package store

import (
	"context"
	"testing"

	"github.com/drunkleen/go-url-shortner/config"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestSaveUrlMapping(t *testing.T) {
	config.LoadConfig()
	// Initialize the StoreService singleton with a mock Redis client.
	_ = InitializeStoreService()
	// Test case 1: Successful mapping storage.
	shortUrl1 := "short-url-1"
	longUrl1 := "long-url-1"
	userId1 := "user-id-1"
	err := SaveUrlMapping(shortUrl1, longUrl1, userId1)
	assert.NoError(t, err)

	// Test case 3: Empty short URL.
	shortUrl2 := ""
	longUrl2 := "long-url-2"
	userId2 := "user-id-2"
	err = SaveUrlMapping(shortUrl2, longUrl2, userId2)
	assert.NoError(t, err)

	// Test case 4: Empty long URL.
	shortUrl3 := "short-url-4"
	longUrl3 := ""
	userId3 := "user-id-"
	err = SaveUrlMapping(shortUrl3, longUrl3, userId3)
	assert.NoError(t, err)

	// Test case 5: Empty user ID.
	shortUrl4 := "short-url-4"
	longUrl4 := "long-url-4"
	userId4 := ""
	err = SaveUrlMapping(shortUrl4, longUrl4, userId4)
	assert.NoError(t, err)
}

func TestRetrieveInitialUrl(t *testing.T) {

	storeService.redisClient = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisURL + ":" + config.AppConfig.RedisPort,
		Password: config.AppConfig.RedisPassword,
		DB:       0,
	})

	// Test case 1: Existing short URL
	shortUrl := "existing-short-url"
	longUrl := "https://example.com"
	err := storeService.redisClient.Set(context.Background(), shortUrl, longUrl, 0).Err()
	assert.NoError(t, err)
	result := RetrieveInitialUrl(shortUrl)
	assert.Equal(t, longUrl, result)

	// Test case 2: Non-existent short URL
	shortUrl = "non-existent-short-url"
	result = RetrieveInitialUrl(shortUrl)
	assert.Empty(t, result)

	// Test case 3: Nil error but empty result
	shortUrl = "nil-error-short-url"
	err = storeService.redisClient.Set(context.Background(), shortUrl, "", 0).Err()
	assert.NoError(t, err)
	result = RetrieveInitialUrl(shortUrl)
	assert.Empty(t, result)
}
