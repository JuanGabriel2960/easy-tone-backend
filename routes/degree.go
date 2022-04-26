package routes

import (
	controllers "github.com/JuanGabriel2960/easy-tone-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func DegreeRouter(route fiber.Router) {
	route.Get("/theory/:degree", controllers.GetTheory)
	route.Get("/exercise/:degree", controllers.GetExercise)
	route.Get("/piece/:degree", controllers.GetPiece)
}
