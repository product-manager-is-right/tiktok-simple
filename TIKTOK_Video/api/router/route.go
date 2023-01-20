package router

import (
	"TIKTOK_Video/api/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister
// 注册路由
func GeneratedRegister(r *server.Hertz) {
	apiRouter := r.Group("/douyin")

	apiRouter.GET("/feed/", handler.Feed)
	apiRouter.POST("/comment/action/", handler.CommentAction)
	apiRouter.GET("/comment/list/", handler.CommentList)
}
