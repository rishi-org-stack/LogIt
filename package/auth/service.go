package auth

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	DB interface {
		FindOrInsert(ctx context.Context, atr *AuthRequest) (interface{}, error)
		Update(ctx context.Context, atr *AuthRequest) (interface{}, error)
		InsertUser(ctx context.Context, atr *AuthRequest) (interface{}, error)
		GetRequest(ctx context.Context, id primitive.ObjectID) (*AuthRequest, error)
	}

	Service interface {
		HandleAuth(ctx context.Context) (*AuthResponse, error)
		GetRequestByID(ctx context.Context, id string) (*AuthRequest, error)
	}
	AuthRequest struct {
		ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Email    string             `json:"email" bson:"email,omitempty"`
		Password string             `json:"password" bson:"password,$default=yet to discuss"`
		Status   string             `json:"status" bson:"status,omitempty"`
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
