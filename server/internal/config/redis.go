package config

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	reidsAddr := os.Getenv("reidsAddr")
	if len(reidsAddr) <= 0 {
		panic("redisAddr not set")
	}
	passWord := os.Getenv("redisPassword")
	if len(passWord) <= 0 {
		panic("redisPassword not set")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     reidsAddr,
		Password: passWord, // no password set
		DB:       0,        // use default DB
	})
	return rdb
}
