package handler

import (
	"context"
	"eshop_user/database"
	"eshop_user/kitex_gen/eshop/user_info"
)

type UserServiceImpl struct {
}

func (u UserServiceImpl) GetUser(ctx context.Context, request *user_info.UserRequest) (r *user_info.UserResponse, err error) {
	user, err := database.QueryOneUserById(nil, request.UserId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, err
	}
	r = &user_info.UserResponse{
		Uid:      user.UID,
		Name:     user.Name,
		Phone:    &user.Phone,
		Email:    &user.Email,
		Password: user.Password,
		Role:     int32(user.Role),
	}
	return r, nil
}
