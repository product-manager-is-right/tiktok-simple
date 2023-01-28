package main

import (
	"TIKTOK_User/api/router"
	"TIKTOK_User/dal"
	"TIKTOK_User/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	initDeps()

	r := server.Default(server.WithHostPorts(":8080"))

	// 注册路由
	router.GeneratedRegister(r)

	r.Spin()
}

func initDeps() {
	// 初始化数据库
	dal.Init()

	// 初始化Jwt
	mw.InitJwt()
}
