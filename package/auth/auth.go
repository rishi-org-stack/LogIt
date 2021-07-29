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

func (authSer AuthService) HandleAuth(ctx context.Context) (interface{}, error) {
	atr := &AuthRequest{
		Email:    "okkkkk mai",
		Password: "password",
	}
	res, err := authSer.AuthData.FindOrInsert(ctx, atr)
	if err != nil {
		return &AuthRequest{}, err
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
		return nil, fmt.Errorf("password and email combination doesn't matched")
	case primitive.ObjectID:
		atr.Status = string(Verified)
		atr.ID = res.(primitive.ObjectID)
		authSer.AuthData.Update(ctx, atr)
		return "success user saved", nil
	}
	return "nil from service", nil
}
