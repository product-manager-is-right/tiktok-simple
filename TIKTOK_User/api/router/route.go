package router

import (
	"GoProject/api/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister
// 注册路由
func GeneratedRegister(r *server.Hertz) {
	apiRouter := r.Group("/douyin/user")

	apiRouter.GET("/", handler.UserInfo)
	apiRouter.POST("/register/", handler.Register)
	apiRouter.POST("/login/", handler.Login)
}
