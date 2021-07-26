package auth

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	DB interface {
		FindOrInsert(ctx context.Context, atr *AuthRequest) (interface{}, error)
	}

	Service interface {
		HandleAuth(ctx context.Context) (interface{}, error)
	}
	AuthRequest struct {
		ID       primitive.ObjectID `bson:"_id,omitempty"`
		Email    string             `bson:"email,omitempty"`
		Password string             `bson:"password,omitempty"`
	}
)
