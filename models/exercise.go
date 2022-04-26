package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exercise struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Exercises []Exercises        `json:"exercises"`
	Degree    string             `json:"degree"`
}

type Exercises struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
