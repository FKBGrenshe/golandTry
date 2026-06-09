package main

import (
	"fmt"

	"git.woa.com/trpc-go/trpc-database/goredis/v3/redisex"

	"git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/log"
)

func main() {
	_ = trpc.NewServer()
	redisExclient, err := redisex.New("trpc.example.test.redis")
	if err != nil {
		log.Fatalf("redis init failed: %v", err)
	}

	client := redisExclient.Client
	v1 := "v1"
	ctx := trpc.BackgroundContext()
	key := "key_redis_universal_v1"

	err = client.Set(ctx, key, v1, 0).Err()
	if err != nil {
		log.Fatalf("set failed: %v", err)
	}

	v, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("get failed: %v", err)
	}
	fmt.Println(v)
}
