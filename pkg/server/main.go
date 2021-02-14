package server

import (
	"fmt"
	"github.com/aliereno/go-redis-docker-learning/pkg/db"
	"github.com/aliereno/go-redis-docker-learning/pkg/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	redisClient := redis.Connect()
	db.InitValues(redisClient)

	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/:name", func(c *fiber.Ctx) error {
		var result db.Country
		err := redis.Get(redisClient, c.Params("name"), &result)
		if err != nil {
			fmt.Errorf(err.Error())
			return err
		}
		return c.JSON(result)
	})

	app.Listen(":3000")
}
