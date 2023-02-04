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

	// 服务间api，获取视频信息
	apiRouter.POST("/publish/GetVideos/", handler.GetVideosByIds)
	apiRouter.POST("/video/favoriteAction/", handler.FavoriteAction)
	// 对外api
	apiRouter.GET("/feed/", handler.Feed, mw.JwtMiddleware.MiddlewareFunc())
	apiRouter.POST("/comment/action/", mw.JwtMiddleware.MiddlewareFunc(), handler.CommentAction)
	apiRouter.GET("/comment/list/", mw.JwtMiddleware.MiddlewareFunc(), handler.CommentList)
	apiRouter.POST("/publish/action/", mw.JwtMiddleware.MiddlewareFunc(), handler.PublishAction)
}
