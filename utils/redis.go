package utils

import (
	"context"
	"92HW/config"
)

func AddTrustedOrigin(userID, origin string) error {
	return config.RedisClient.SAdd(context.Background(), userID, origin).Err()
}

func RemoveTrustedOrigin(userID, origin string) error {
	return config.RedisClient.SRem(context.Background(), userID, origin).Err()
}

func GetTrustedOrigins(userID string) ([]string, error) {
	return config.RedisClient.SMembers(context.Background(), userID).Result()
}
