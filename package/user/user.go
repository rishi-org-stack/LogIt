package user

import (
	"context"
)

type (
	UserService struct {
		UserData DB
	}
)

func Init(db DB) *UserService {
	return &UserService{
		UserData: db,
	}
}
func (uSer UserService) GetUser(ctx context.Context, id string) ([]UserAggregate, error) {
	userRes, err := uSer.UserData.GetUser(ctx, id)
	if err != nil {
		return []UserAggregate{}, err
	}
	return userRes, nil
}
func (uSer UserService) UpdateUser(ctx context.Context, user *User) (*User, error) {
	_, err := uSer.UserData.UpdateUser(ctx, user)
	if err != nil {
		return &User{}, err
	}
	return &User{}, nil
}
