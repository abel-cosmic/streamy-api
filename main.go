package main

import (
	connector "github.com/abel-cosmic/streamy-api/database"
	"github.com/abel-cosmic/streamy-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
)

func main() {
	fastergoding.Run()
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 2048 * 2048,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Welcome to Streamy API version 1.0",
		})
	})
	db := connector.Connect()
	// Attach routes to API group
	router.Routes(app, db)
	app.Listen(":3000")
}
