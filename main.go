package main

import (
	"fmt"
	"os"

	env "github.com/JuanGabriel2960/easy-tone-backend/env"
	routes "github.com/JuanGabriel2960/easy-tone-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	env.Load()

	app := fiber.New()
	app.Use(cors.New())

	routes.DegreeRouter(app.Group("api/degree"))
	routes.ResourcesRouter(app.Group("api/resources"))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("easy-tone-backend")
	})

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
