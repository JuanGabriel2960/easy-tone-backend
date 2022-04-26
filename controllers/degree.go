package controllers

import (
	"context"

	database "github.com/JuanGabriel2960/easy-tone-backend/database"
	models "github.com/JuanGabriel2960/easy-tone-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.Background()

func GetTheory(c *fiber.Ctx) error {
	degree := c.Params("degree")

	var collection = database.GetCollection("theories")
	var theory models.Theory

	filter := bson.M{"degree": degree}

	cur := collection.FindOne(ctx, filter)

	err := cur.Decode(&theory)

	if err != nil {
		return err
	}

	return c.JSON(theory)
}

func GetExercise(c *fiber.Ctx) error {
	degree := c.Params("degree")

	var collection = database.GetCollection("exercises")
	var exercise models.Exercise

	filter := bson.M{"degree": degree}

	cur := collection.FindOne(ctx, filter)

	err := cur.Decode(&exercise)

	if err != nil {
		return err
	}

	return c.JSON(exercise)
}
