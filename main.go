package main

import (
	connector "github.com/abel-cosmic/streamy-api/database"
	"github.com/abel-cosmic/streamy-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := connector.Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Starter API for golang fiber ",
		})
	})

	router.Routes(app, db)
	app.Listen(":3000")
}
