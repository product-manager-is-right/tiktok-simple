package router

import (
	"TIKTOK_Video/api/handler"
	"TIKTOK_Video/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister
// 注册路由
func GeneratedRegister(r *server.Hertz) {
	apiRouter := r.Group("/douyin")

	apiRouter.GET("/feed/", handler.Feed, mw.JwtMiddleware.MiddlewareFunc())
	apiRouter.POST("/comment/action/", mw.JwtMiddleware.MiddlewareFunc(), handler.CommentAction)
	apiRouter.GET("/comment/list/", mw.JwtMiddleware.MiddlewareFunc(), handler.CommentList)

	//publish action
	apiRouter.POST("/publish/action/", handler.PublishAction)
}
