package main

import (
	"eshop_user/handler"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	group := r.Group("/api/v1/user")
	group.POST("/login", handler.HandleLogin)
}
