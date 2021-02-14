package redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"os"
)

func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		Password: "",
		DB:       0,
	})
	return client
}
func Set(c *redis.Client, key string, value interface{}) error {
	p, err := json.Marshal(value)
	if err != nil {
		fmt.Errorf(err.Error())
		return err
	}
	return c.Set(key, p, 0).Err()
}
func Get(c *redis.Client, key string, dest interface{}) error {
	result, err := c.Get(key).Result()
	if err == redis.Nil {
		var err = fiber.Error{
			Code:    404,
			Message: "Country not found.",
		}
		return &err
	} else if err != nil {
		fmt.Errorf(err.Error())
		return err
	}
	return json.Unmarshal([]byte(result), dest)
}
