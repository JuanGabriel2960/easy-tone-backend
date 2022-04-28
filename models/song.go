package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Song struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `json:"name"`
	Author string             `json:"author"`
	Url    string             `json:"url"`
	Image  string             `json:"image"`
}

type Songs []*Song
