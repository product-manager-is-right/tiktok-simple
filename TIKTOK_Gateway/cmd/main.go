package main

import (
	"TIKTOK_Gateway/configs"
	"TIKTOK_Gateway/route"
	"github.com/cloudwego/hertz/pkg/app/server"
	"log"
	"strconv"
)

func main() {
	// 读取config配置
	myConfig, err := configs.ReadConfig()
	if err != nil {
		log.Fatal("文件读取失败:", err.Error())
	}

	// 创建服务
	addr := ":" + strconv.Itoa(myConfig.Port)
	h := server.New(server.WithHostPorts(addr))

	// 路由注册
	route.Register(myConfig, h)

	// 启动服务
	h.Spin()
}
