package handler

import (
	"eshop_user/common/constant"
	"eshop_user/database"
	"eshop_user/model"
	"eshop_user/model/req"
	"eshop_user/model/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func HandleLogin(ctx *gin.Context) {
	var (
		param req.LoginReqDTO
		res   resp.LoginRespDTO
	)
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusInternalServerError, "参数绑定失败")
		return
	}
	user, err := database.QueryOneUserByName(nil, param.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "服务器内部错误")
		return
	}
	if user.Password != param.Password {
		res = resp.LoginRespDTO{
			Code: constant.WrongPassword,
			Info: "密码错误",
		}
		ctx.JSON(http.StatusOK, res)
		return
	}
	res = resp.LoginRespDTO{
		Code: constant.Success,
		Info: "success",
		Uid:  user.UID,
	}
	ctx.JSON(http.StatusOK, res)
	return
}

func HandleRegister(ctx *gin.Context) {
	var (
		param req.RegisterReqDTO
		res   resp.RegisterRespDTO
	)
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusInternalServerError, "参数绑定失败")
		return
	}
	user, err := database.QueryOneUserByName(nil, param.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "服务器内部错误")
		return
	}
	if user != nil {
		res = resp.RegisterRespDTO{
			Code: constant.DuplicatedUser,
			Info: "用户已存在",
		}
		ctx.JSON(http.StatusOK, res)
		return
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			fmt.Printf("生成UUID失败：%v", err)
			return
		}
		user := &model.User{
			UID:        id.String(),
			Name:       param.Name,
			Phone:      "",
			Email:      "",
			Password:   param.Password,
			Role:       0,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			IsDeleted:  false,
		}
		err = database.InsertOneUser(nil, user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, "服务器内部错误")
			return
		}
		res = resp.RegisterRespDTO{
			Code: constant.Success,
			Info: "success",
			Uid:  user.UID,
		}
		ctx.JSON(http.StatusOK, res)
		return
	}
}

func HandleGetUser(ctx *gin.Context) {
	var (
		uid string
		res resp.UserRespDTO
	)
	uid = ctx.Param("uid")
	user, err := database.QueryOneUserById(nil, uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "服务器内部错误")
		return
	}
	if user == nil {
		res = resp.UserRespDTO{
			Code: constant.UserNotFound,
			Info: "用户不存在",
			User: nil,
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	res = resp.UserRespDTO{
		Code: constant.Success,
		Info: "success",
		User: user,
	}
	ctx.JSON(http.StatusOK, res)
	return
}
