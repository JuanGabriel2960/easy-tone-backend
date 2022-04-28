package controllers

import (
	"context"

	database "github.com/JuanGabriel2960/easy-tone-backend/database"
	models "github.com/JuanGabriel2960/easy-tone-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var ctxx = context.Background()

func GetSongs(c *fiber.Ctx) error {
	var collection = database.GetCollection("songs")
	var songs models.Songs

	filter := bson.D{}

	cur, err := collection.Find(ctxx, filter)

	if err != nil {
		return err
	}

	for cur.Next(ctxx) {
		var song models.Song
		err = cur.Decode(&song)

		if err != nil {
			return err
		}

		songs = append(songs, &song)
	}

	return c.JSON(songs)
}
