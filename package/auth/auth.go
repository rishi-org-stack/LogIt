package auth

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// import agm "logit/v1/package/auth/databases/mongo"

type AuthService struct {
	AuthData DB
	JwtSer   TokenGenratorInterface
}

func Init(db DB, js TokenGenratorInterface) *AuthService {
	return &AuthService{
		AuthData: db,
		JwtSer:   js,
	}
}

func (authSer AuthService) HandleAuth(ctx context.Context) (*AuthResponse, error) {
	atr := &AuthRequest{
		Email:    "okkkkk mai",
		Password: "password",
	}
	res, err := authSer.AuthData.FindOrInsert(ctx, atr)
	if err != nil {
		return &AuthResponse{}, err
	}
	switch res.(type) {
	case *AuthRequest:
		resA := res.(*AuthRequest)
		if resA.Password == atr.Password {
			token, err := authSer.JwtSer.GenrateToken(resA.ID.String(), resA.Email)
			if err != nil {
				return nil, err
			}
			return &AuthResponse{
				Token: token,
			}, nil
		}
		return &AuthResponse{}, fmt.Errorf("password and email combination doesn't matched")
	case primitive.ObjectID:
		atr.Status = string(Verified)
		atr.ID = res.(primitive.ObjectID)
		_, err := authSer.AuthData.Update(ctx, atr)
		if err != nil {
			return &AuthResponse{}, err
		}
		// atr.ID = updateID.(primitive.ObjectID)
		_, err = authSer.AuthData.InsertUser(ctx, atr)
		if err != nil {
			return &AuthResponse{}, err
		}
		token, err := authSer.JwtSer.GenrateToken(atr.ID.String(), atr.Email)
		if err != nil {
			return nil, err
		}
		return &AuthResponse{
			Token: token,
		}, nil
	}
	return &AuthResponse{}, nil
}
