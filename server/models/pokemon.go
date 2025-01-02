package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pokemon struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name            string             `bson:"name" json:"name"`
	Number          int                `bson:"number" json:"number"`
	ImageURL        string             `bson:"imageUrl" json:"imageUrl"`
	Description     string             `bson:"description" json:"description"`
	BackgroundColor string             `bson:"backgroundColor" json:"backgroundColor"`
}
