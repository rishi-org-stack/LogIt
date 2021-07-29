package auth

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	DB interface {
		FindOrInsert(ctx context.Context, atr *AuthRequest) (interface{}, error)
		Update(ctx context.Context, atr *AuthRequest) (interface{}, error)
		InsertUser(ctx context.Context, atr *AuthRequest) (interface{}, error)
	}

	Service interface {
		HandleAuth(ctx context.Context) (*AuthResponse, error)
	}
	AuthRequest struct {
		ID       primitive.ObjectID `bson:"_id,omitempty"`
		Email    string             `bson:"email,omitempty"`
		Password string             `bson:"password,$default=yet to discuss"`
		Status   string             `bson:"status,omitempty"`
	}
	StatusType string

	TokenGenratorInterface interface {
		GenrateToken(id, email string) (string, error)
	}
	//DTO's
	AuthResponse struct {
		Token string `json:"token"`
	}
	// AuthRequest struct {
	// 	ID       primitive.ObjectID `json:"id,omitempty"`
	// 	Email    string             `json:"email,omitempty"`
	// 	Password string             `json:"password"`
	// }
)

const (
	New      StatusType = "New"
	Verified StatusType = "Verified"
	Invalid  StatusType = "Invalid"
	Old      StatusType = "Old"
)
