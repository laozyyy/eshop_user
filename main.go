package main

import (
	"eshop_user/database"
	"eshop_user/kitex_gen/eshop/user_info/userservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)

func main() {
	database.Init()
	svr := userservice.NewServer(
		new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "UserService"}),
	)
	err := svr.Run()
	if err != nil {
		panic(err)
	}

}

//package main
//
//import (
//	"fmt"
//	"github.com/bytedance/sonic"
//	"github.com/go-zookeeper/zk"
//	"log"
//	"time"
//)
//
//// 定义配置结构体，这里假设配置包含服务器地址、端口和超时时间，可根据实际需求修改
//type Config struct {
//	ID   string `json:"tag_id"`
//	Name string `json:"tag_name"`
//}
//
//func main() {
//	// Zookeeper服务器地址列表，可配置多个地址以实现高可用，这里示例只写一个
//	hosts := []string{"117.72.72.114:2181"}
//	// 配置数据在Zookeeper中的节点路径，需替换为实际存储配置的路径
//	configNodePath := "/frontend/menu"
//
//	// 创建Zookeeper连接
//	conn, _, err := zk.Connect(hosts, time.Second*5)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer conn.Close()
//
//	// 获取配置节点的数据
//	data, _, err := conn.Get(configNodePath)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 将获取到的数据（字节切片类型）解析为Config结构体
//	var config []*Config
//	err = sonic.Unmarshal(data, &config)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, t := range config {
//		fmt.Println(t.Name)
//		fmt.Println(t.ID)
//	}
//
//}
