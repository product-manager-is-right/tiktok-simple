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
	//微服务
	publishRouter.POST("/UserVideo/", handler.PublishVideo)
	publishRouter.Use(mw.JwtMiddleware.MiddlewareFunc())
	publishRouter.GET("/list/", handler.PublishList)

	// favorite路由组
	favoriteRouter := r.Group("/douyin/favorite")
	//微服务接口
	favoriteRouter.GET("/IsFavor/", handler.IsFavorite)
	favoriteRouter.Use(mw.JwtMiddleware.MiddlewareFunc())
	favoriteRouter.POST("/action/", handler.FavoriteAction)
	favoriteRouter.GET("/list/", handler.FavoriteList)

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
