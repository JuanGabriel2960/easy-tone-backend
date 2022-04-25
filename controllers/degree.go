package controllers

import (
	"context"

	database "github.com/JuanGabriel2960/easy-tone-backend/database"
	models "github.com/JuanGabriel2960/easy-tone-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.GetCollection("theories")
var ctx = context.Background()

func GetTheory(c *fiber.Ctx) error {
	degree := c.Params("degree")

	var theory models.Theory

	filter := bson.M{"degree": degree}

	cur := collection.FindOne(ctx, filter)

	err := cur.Decode(&theory)

	if err != nil {
		return err
	}

	return c.JSON(theory)
}
