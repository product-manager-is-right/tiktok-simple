package router

import (
	"TIKTOK_User/api/handler"
	"TIKTOK_User/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister
// 注册路由
func GeneratedRegister(r *server.Hertz) {

	// user路由组
	userRouter := r.Group("/douyin/user")
	// 服务间api
	userRouter.POST("/get", handler.UserInfoList)
	// 对外api
	userRouter.GET("/", mw.JwtMiddleware.MiddlewareFunc(), handler.UserInfo)
	userRouter.POST("/register/", handler.Register, mw.JwtMiddleware.LoginHandler)
	userRouter.POST("/login/", mw.JwtMiddleware.LoginHandler)

	// publish路由组
	publishRouter := r.Group("/douyin/publish")
	// 服务间api
	publishRouter.POST("/UserVideo/", handler.PublishVideo)
	// 对外api
	publishRouter.Use(mw.JwtMiddleware.MiddlewareFunc())
	publishRouter.GET("/list/", handler.PublishList)

	// favorite路由组
	favoriteRouter := r.Group("/douyin/favorite")
	// 服务间api
	favoriteRouter.GET("/IsFavor/", handler.IsFavorite)
	// 对外api
	favoriteRouter.Use(mw.JwtMiddleware.MiddlewareFunc())
	favoriteRouter.POST("/action/", handler.FavoriteAction)
	favoriteRouter.GET("/list/", handler.FavoriteList)

	// relation路由组
	relationRouter := r.Group("/douyin/relation")
	// 对外api
	relationRouter.Use(mw.JwtMiddleware.MiddlewareFunc())
	relationRouter.POST("/action/", handler.RelationAction)
	relationRouter.GET("/follow/list/", handler.FollowList)
	relationRouter.GET("/follower/list/", handler.FollowerList)
	relationRouter.GET("/friend/list/", handler.FriendList)
}
