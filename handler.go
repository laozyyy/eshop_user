package main

import (
	"context"
	"eshop_user/common/constant"
	"eshop_user/database"
	"eshop_user/kitex_gen/eshop/user_info"
	"eshop_user/model"
	"time"
)

type UserServiceImpl struct {
}

func (u UserServiceImpl) GetOneUserByName(ctx context.Context, name string) (r *user_info.GetOneUserResponse, err error) {
	user, err := database.QueryOneUserByName(nil, name)
	if err != nil {
		r = &user_info.GetOneUserResponse{
			User:   nil,
			Code:   constant.InternalServerError,
			ErrStr: constant.CodeInfo(constant.InternalServerError),
		}
		return r, err
	}
	if user == nil {
		r = &user_info.GetOneUserResponse{
			User:   nil,
			Code:   constant.UserNotFound,
			ErrStr: constant.CodeInfo(constant.UserNotFound),
		}
		return r, err
	}
	uuser := &user_info.User{
		Uid:      user.UID,
		Name:     user.Name,
		Phone:    &user.Phone,
		Email:    &user.Email,
		Password: user.Password,
		Role:     int32(user.Role),
	}
	r = &user_info.GetOneUserResponse{
		User: uuser,
		Code: constant.Success,
	}
	return r, nil
}

func (u UserServiceImpl) GetOneUser(ctx context.Context, request *user_info.GetOneUserRequest) (r *user_info.GetOneUserResponse, err error) {
	user, err := database.QueryOneUserById(nil, request.Uid)
	if err != nil {
		r = &user_info.GetOneUserResponse{
			User:   nil,
			Code:   constant.InternalServerError,
			ErrStr: constant.CodeInfo(constant.InternalServerError),
		}
		return r, err
	}
	if user == nil {
		r = &user_info.GetOneUserResponse{
			User:   nil,
			Code:   constant.UserNotFound,
			ErrStr: constant.CodeInfo(constant.UserNotFound),
		}
		return r, err
	}
	uuser := &user_info.User{
		Uid:      user.UID,
		Name:     user.Name,
		Phone:    &user.Phone,
		Email:    &user.Email,
		Password: user.Password,
		Role:     int32(user.Role),
	}
	r = &user_info.GetOneUserResponse{
		User: uuser,
		Code: constant.Success,
	}
	return r, nil
}

func (u UserServiceImpl) InsertOneUser(ctx context.Context, user *user_info.User) (r *user_info.InsertOneUserResponse, err error) {
	uuser := &model.User{
		UID:        user.Uid,
		Name:       user.Name,
		Phone:      *user.Phone,
		Email:      *user.Email,
		Password:   user.Password,
		Role:       0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		IsDeleted:  false,
	}
	err = database.InsertOneUser(nil, uuser)
	if err != nil {
		r = &user_info.InsertOneUserResponse{
			Code:   constant.InternalServerError,
			ErrStr: constant.CodeInfo(constant.InternalServerError),
		}
		return r, err
	}
	r = &user_info.InsertOneUserResponse{
		Code: constant.Success,
	}
	return
}
