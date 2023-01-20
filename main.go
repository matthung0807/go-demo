package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := NewRedisClient()
	err := rdb.Set(ctx, "name", "John", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val) // John
}

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "12345",
		DB:       0, // use default DB
	})
}
