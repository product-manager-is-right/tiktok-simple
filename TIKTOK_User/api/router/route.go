package router

import (
	"GoProject/api/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister
// 注册路由
func GeneratedRegister(r *server.Hertz) {
	// user路由组
	userRouter := r.Group("/douyin/user")
	userRouter.GET("/", handler.UserInfo)
	userRouter.POST("/register/", handler.Register)
	userRouter.POST("/login/", handler.Login)

	// publish路由组
	publishRouter := r.Group("/douyin/publish")
	publishRouter.POST("/action/", handler.PublishAction)
	publishRouter.GET("/list/", handler.PublishList)

	// favorite路由组
	favoriteRouter := r.Group("/douyin/favorite")
	favoriteRouter.POST("/action/", handler.FavoriteAction)
	favoriteRouter.GET("/list/", handler.FavoriteList)

	// relation路由组
	relationRouter := r.Group("/douyin/relation")
	relationRouter.POST("/action/", handler.RelationAction)
	relationRouter.GET("/follow/list/", handler.FollowList)
	relationRouter.GET("/follower/list/", handler.FollowerList)
	relationRouter.GET("/friend/list/", handler.FriendList)
}
