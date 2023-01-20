package main

import (
	"TIKTOK_Video/api/router"
	"TIKTOK_Video/dal"
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
	// 初始化数据库
	dal.Init()
}
