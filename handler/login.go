package handler

import (
	"eshop_user/constant"
	"eshop_user/database"
	"eshop_user/model/req"
	"eshop_user/model/resp"
	"github.com/gin-gonic/gin"
	"net/http"
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
			Code: constant.Success, //todo
			Info: "密码错误",
		}
		ctx.JSON(http.StatusOK, res)
		return
	}
	res = resp.LoginRespDTO{
		Code: constant.Success, //todo
		Info: "success",
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
			Code: constant.DuplicatedUser, //todo
			Info: "用户已存在",
		}
		ctx.JSON(http.StatusOK, res)
		return
	} else {
		res = resp.RegisterRespDTO{
			Code: constant.Success, //todo
			Info: "success",
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
	if err := ctx.ShouldBindQuery(&uid); err != nil {
		ctx.JSON(http.StatusInternalServerError, "参数绑定失败")
		return
	}
	user, err := database.QueryOneUserById(nil, uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "服务器内部错误")
		return
	}
	if user == nil {
		res = resp.UserRespDTO{
			Code: constant.UserNotFound, //todo
			Info: "用户不存在",
			User: nil,
		}
		ctx.JSON(http.StatusOK, res)
		return
	}

	res = resp.UserRespDTO{
		Code: constant.Success, //todo
		Info: "success",
		User: user,
	}
	ctx.JSON(http.StatusOK, res)
	return
}
