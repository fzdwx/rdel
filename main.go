package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
)

func main() {
	host := flag.String("host", "localhost", "redis host")
	port := flag.String("port", "6379", "redis port")
	password := flag.String("password", "", "redis password")
	db := flag.Int("db", 0, "redis db")
	delKeyPrefix := flag.String("key", "", "redis del key prefix")

	flag.Parse()
	if *delKeyPrefix == "" {
		return
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", *host, *port),
		Password: *password,
		DB:       *db,
	})

	ctx := context.Background()
	keys, err := rdb.Keys(ctx, *delKeyPrefix).Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	chunks := lo.Chunk(keys, 100)
	for _, chunk := range chunks {
		rdb.Del(ctx, chunk...)
	}
}
