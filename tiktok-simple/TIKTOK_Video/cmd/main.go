package main

import (
	"TIKTOK_Video/api/router"
	"TIKTOK_Video/dal"
	"TIKTOK_Video/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	initDeps()

	r := server.Default(server.WithHostPorts(":8081"))

	// 注册路由
	router.GeneratedRegister(r)

	r.Spin()
}

func initDeps() {
	// 初始化数据库
	dal.Init()

	// 初始化jwt
	mw.InitJwt()
}
