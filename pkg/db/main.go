package db

import "github.com/go-redis/redis"

func InitValues(client *redis.Client) {
	InitCountry(client)
}
