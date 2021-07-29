package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Idea struct {
		ID          primitive.ObjectID `bson:"_id,omitempty"`
		Name        string             `bson:"name,omitempty"`
		Description string             `bson:"description,default=yet to dicuss"`
	}
)
