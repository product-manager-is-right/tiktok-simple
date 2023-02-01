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
	userRouter.GET("/", mw.JwtMiddleware.MiddlewareFunc(), handler.UserInfo)
	userRouter.POST("/register/", handler.Register, mw.JwtMiddleware.LoginHandler)
	userRouter.POST("/login/", mw.JwtMiddleware.LoginHandler)

	// publish路由组
	publishRouter := r.Group("/douyin/publish")

	publishRouter.Use(mw.JwtMiddleware.MiddlewareFunc())
	publishRouter.GET("/list/", handler.PublishList)
	publishRouter.POST("/UserVideo/", handler.PublishVideo)

	// favorite路由组
	favoriteRouter := r.Group("/douyin/favorite")

	favoriteRouter.Use(mw.JwtMiddleware.MiddlewareFunc())

	favoriteRouter.POST("/action/", handler.FavoriteAction)
	favoriteRouter.GET("/list/", handler.FavoriteList)
	favoriteRouter.GET("/IsFavor/", handler.IsFavorite)

	// relation路由组
	relationRouter := r.Group("/douyin/relation")

	relationRouter.Use(mw.JwtMiddleware.MiddlewareFunc())

	relationRouter.POST("/action/", handler.RelationAction)
	relationRouter.GET("/follow/list/", handler.FollowList)
	relationRouter.GET("/follower/list/", handler.FollowerList)
	relationRouter.GET("/friend/list/", handler.FriendList)


	//添加一个由微服务之间调用的请求路径,获取用户信息列表,额外开一个是为了不走JWT认证
	r.POST("douyin/user/get", handler.UserInfoList)

}
