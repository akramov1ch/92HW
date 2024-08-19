package config

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func SetupConfig() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := os.Getenv("REDIS_DB")
	redisAddr := redisHost + ":" + redisPort
	redisDBInt, err := strconv.Atoi(redisDB)
	if err != nil {
		log.Fatal(err)
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDBInt,
	})
}

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	ctx := context.Background()

	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Redisga ulanishda xatolik: %v", err)
	}
}
