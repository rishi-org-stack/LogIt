package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	User struct {
		ID     primitive.ObjectID             `bson:"_id,omitempty"`
		Name   string                         `bson:"name"`
		AuthID primitive.ObjectID             `bson:"auth_id"`
		Ideas  map[primitive.ObjectID]*Status `bson:"projects"`
	}
	Status struct {
		MarkedAs  string `bson:"marked_as"`
		Deadline  string `bson:"dealine"`
		AccessKey []byte `bson:"access_key"`
	}
)
