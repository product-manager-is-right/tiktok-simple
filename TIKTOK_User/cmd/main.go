package main

import (
	"GoProject/api/router"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	initDeps()

	r := server.Default()

	// 注册路由
	router.GeneratedRegister(r)

	r.Spin()
}

func initDeps() {
	// 初始化数据库等
}