package main

import (
	routes "github.com/JuanGabriel2960/easy-tone-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	routes.DegreeRouter(app.Group("api/degree"))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("easy-tone-backend")
	})

	app.Listen(":3000")
}
