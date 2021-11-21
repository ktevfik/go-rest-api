package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mtevfik41/alfa/database"
	"github.com/mtevfik41/alfa/router"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("API is running")
		return err
	})

	// Setup the router
	router.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":3000")
}
