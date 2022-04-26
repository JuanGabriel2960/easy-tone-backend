package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Piece struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Pieces []string           `json:"pieces"`
	Degree string             `json:"degree"`
}
