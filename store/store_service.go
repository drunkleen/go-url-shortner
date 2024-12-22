package store

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	storeService = &StoreService{}
	ctx          = context.Background()
)

const CacheDuration = 24 * time.Hour

type StoreService struct {
	redisClient *redis.Client
}

func InitializeStoreService() *StoreService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Panicf("Failed to connect to Redis: %v", err)
	}

	log.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortUrl, longUrl, userId string) {
	if err := storeService.redisClient.Set(ctx, shortUrl, longUrl, CacheDuration).Err(); err != nil {
		log.Printf("Failed to save url mapping: %v", err)
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		log.Printf("Failed to retrieve initial url: %v", err)
	}
	return result
}
