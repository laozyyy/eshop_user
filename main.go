package main

import (
	"eshop_user/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	initRouter(r)
	_ = r.Run(":9001")
}
