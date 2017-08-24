package main

import (
	"github.com/go-redis/redis"
)

var redisClient = newRedisClient()

func newRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     *raddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	log.Println(pong, err)
	return client
}
