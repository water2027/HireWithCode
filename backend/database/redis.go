package database

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
	rdb *redis.Client
}

var (
	redisClient *RedisClient = &RedisClient{
		rdb: nil,
	}
	redisOnce     sync.Once
)

func GetRedisClient() *RedisClient {
	redisOnce.Do(func() {
		if redisClient == nil || redisClient.rdb == nil {
			redisClient = &RedisClient{}
			initRedis()
		}
	})
	return redisClient
}

func initRedis() {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379"
	}
	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Println("Connecting to Redis at", addr)
	redisClient.rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
}

func (rc *RedisClient) SetValue(key string, value string, expire time.Duration) error {
	if rc.rdb == nil {
		initRedis()
	}
	return rc.rdb.Set(ctx, key, value, expire).Err()
}

func (rc *RedisClient) GetValue(key string) (string, error) {
	if rc.rdb == nil {
		initRedis()
	}
	return rc.rdb.Get(ctx, key).Result()
}

func (rc *RedisClient) DeleteValue(key string) error {
	if rc.rdb == nil {
		initRedis()
	}
	return rc.rdb.Del(ctx, key).Err()
}
