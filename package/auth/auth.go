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

func (authSer AuthService) HandleAuth(ctx context.Context, ar *AuthRequest) (*AuthResponse, error) {
	// atr := &AuthRequest{
	// 	Email:    "rishi@gmail.com",
	// 	Password: "password",
	// }
	res, err := authSer.AuthData.FindOrInsert(ctx, ar)
	if err != nil {
		return &AuthResponse{}, err
	}
	switch res.(type) {
	case *AuthRequest:
		resA := res.(*AuthRequest)
		if resA.Password == ar.Password {
			token, err := authSer.JwtSer.GenrateToken(resA.ID.Hex(), resA.Email)
			if err != nil {
				return nil, err
			}
			return &AuthResponse{
				Token: token,
			}, nil
		}
		return &AuthResponse{}, fmt.Errorf("password and email combination doesn't matched")
	case primitive.ObjectID:
		ar.Status = string(Verified)
		ar.ID = res.(primitive.ObjectID)
		_, err := authSer.AuthData.Update(ctx, ar)
		if err != nil {
			return &AuthResponse{}, err
		}
		// ar.ID = updateID.(primitive.ObjectID)
		_, err = authSer.AuthData.InsertUser(ctx, ar)
		if err != nil {
			return &AuthResponse{}, err
		}
		token, err := authSer.JwtSer.GenrateToken(ar.ID.Hex(), ar.Email)
		if err != nil {
			return nil, err
		}
		return &AuthResponse{
			Token: token,
		}, nil
	}
	return &AuthResponse{}, nil
}

func (ar AuthService) GetRequestByID(ctx context.Context, id string) (*AuthRequest, error) {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &AuthRequest{}, err
	}
	authR, err := ar.AuthData.GetRequest(ctx, Id)
	if err != nil {
		return &AuthRequest{}, nil
	}
	return authR, nil
}
