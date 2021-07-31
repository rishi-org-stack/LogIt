package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	au "logit/v1/package/auth"
)

type (
	DB interface {
		GetUser(ctx context.Context, id string) ([]UserAggregate, error)
		UpdateUser(ctx context.Context, user *User) (*User, error)
	}
	Service interface {
		UpdateUser(ctx context.Context, user *User) (*User, error)
		GetUser(ctx context.Context, id string) ([]UserAggregate, error)
	}
	//TODO:User ID needs to be of type Object Id
	User struct {
		ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Name   string             `bson:"name"`
		AuthID primitive.ObjectID `bson:"auth_id"`
		Ideas  map[string]*Status `bson:"ideas"`
	}
	UserAggregate struct {
		ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Name   string             `bson:"name"`
		AuthID primitive.ObjectID `bson:"auth_id"`
		Ideas  map[string]*Status `bson:"ideas"`
		Auth   au.AuthRequest     `bson:"auth,omitempty"`
	}
	Status struct {
		MarkedAs  string `bson:"marked_as"`
		Deadline  string `bson:"dealine"`
		AccessKey []byte `bson:"access_key"`
	}
)
