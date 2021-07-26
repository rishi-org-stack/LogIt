package auth

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// import agm "logit/v1/package/auth/databases/mongo"

type AuthService struct {
	AuthData DB
}

func Init(db DB) *AuthService {
	return &AuthService{
		AuthData: db,
	}
}

func (authSer AuthService) HandleAuth(ctx context.Context) (interface{}, error) {
	atr := &AuthRequest{
		Email:    "jhaji",
		Password: "password",
	}
	res, err := authSer.AuthData.FindOrInsert(ctx, atr)
	if err != nil {
		return &AuthRequest{}, err
	}
	switch (res).(type) {
	case *AuthRequest:
		resA := res.(*AuthRequest)
		if resA.Password == atr.Password {
			return "success password matched", nil
		}
		return nil, fmt.Errorf("password doesn't matched")
	case primitive.ObjectID:
		return "success user saved", nil
	}
	return "nil from service", nil
}
