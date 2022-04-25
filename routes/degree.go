package routes

import (
	controllers "github.com/JuanGabriel2960/easy-tone-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func DegreeRouter(route fiber.Router) {
	route.Get("/:degree", controllers.GetTheory)
}
