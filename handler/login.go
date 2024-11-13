package handler

import (
	"eshop_user/database"
	"eshop_user/model/req"
	"eshop_user/model/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleLogin(ctx *gin.Context) {
	var (
		param req.LoginDTO
		res   resp.LoginResp
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
		res = resp.LoginResp{
			Code: "123", //todo
			Info: "密码错误",
		}
		ctx.JSON(http.StatusOK, res)
		return
	}
	res = resp.LoginResp{
		Code: "1234", //todo
		Info: "成功",
	}
	ctx.JSON(http.StatusOK, res)
	return
}
