package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Theory struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Data   Data               `json:"data"`
	Degree string             `json:"degree"`
}

type Data struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Theories []*Theory
