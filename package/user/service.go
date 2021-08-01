package user

import (
	"context"
	au "logit/v1/package/auth"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	DB interface {
		GetUser(ctx context.Context, id primitive.ObjectID) (*User, error)
		UpdateUser(ctx context.Context, user *User) (*User, error)
	}
	Service interface {
		UpdateUser(ctx context.Context, user *User) (*User, error)
		GetUser(ctx context.Context, id string) (*UserAggregate, error)
	}
	//TODO:User ID needs to be of type Object Id
	User struct {
		ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Name   string             `json:"name" bson:"name"`
		AuthID primitive.ObjectID `json:"authID" bson:"auth_id"`
		Ideas  map[string]*Status `json:"ideas" bson:"ideas"`
	}
	UserAggregate struct {
		User
		Auth au.AuthRequest `json:"auth"`
	}
	Status struct {
		MarkedAs  string `bson:"marked_as"`
		Deadline  string `bson:"dealine"`
		AccessKey []byte `bson:"access_key"`
	}
)
