package routes

import (
	controllers "github.com/JuanGabriel2960/easy-tone-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func ResourcesRouter(route fiber.Router) {
	route.Get("/songs", controllers.GetSongs)
}
