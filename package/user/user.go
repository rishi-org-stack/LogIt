package user

import (
	"context"
	"fmt"
	"logit/v1/package/auth"
	"logit/v1/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	UserService struct {
		UserData    DB
		AuthService auth.Service
	}
)

func Init(db DB, authser auth.Service) *UserService {
	return &UserService{
		UserData:    db,
		AuthService: authser,
	}
}
func (uSer UserService) GetUser(ctx context.Context, id string) (*UserAggregate, error) {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &UserAggregate{}, err
	}
	userRes, err := uSer.UserData.GetUser(ctx, Id)
	if err != nil {
		return &UserAggregate{}, err
	}
	authID := userRes.AuthID
	authreq, err := uSer.AuthService.GetRequestByID(ctx, authID.Hex())
	if err != nil {
		return &UserAggregate{}, err
	}
	res := &UserAggregate{}
	res.Auth = *authreq
	res.User = *userRes
	if err != nil {
		return &UserAggregate{}, err
	}
	return res, nil
}
func (uSer UserService) UpdateUser(ctx context.Context, user *User) (*User, error) {
	userdbStruct, err := uSer.UserData.GetUser(ctx, user.ID)
	if err != nil {
		return &User{}, err
	}
	_, err = uSer.UserData.UpdateUser(ctx, user)
	if err != nil {
		return &User{}, err
	}

	util.TransferData(*user, *userdbStruct)
	fmt.Println(userdbStruct)
	return &User{}, nil
}
