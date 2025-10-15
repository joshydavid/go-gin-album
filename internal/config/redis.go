package config

import (
	"context"
	"fmt"
	"go-gin-album/pkg/util"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func SetUpRedisClient(ctx context.Context) *redis.Client {
	util.LoadEnv()
	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("REDIS_PORT"))
	password := os.Getenv("REDIS_PASSWORD")
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
		Protocol: 2,
	})

	err := client.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("❌ Could not connect to Redis: %v\n", err)
	}

	log.Println("✅ Successfully connected to Redis.")
	return client
}
