package main

import (
	"github.com/aliereno/go-redis-docker-learning/src/redis"
	"github.com/aliereno/go-redis-docker-learning/src/server"
)

func main() {
	redisClient := redis.Connect()
	_ = redisClient

	server.Start()
}