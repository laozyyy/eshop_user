package main

import (
	"eshop_user/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	initRouter(r)
	_ = r.Run(":9000")
	//addr := "127.0.0.1:8888"
	//svr := userservice.NewServer(new(handler.UserServiceImpl),
	//	server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "UserService"}),
	//)
	//err := svr.Run()
	//if err != nil {
	//	panic(err)
	//}
}
