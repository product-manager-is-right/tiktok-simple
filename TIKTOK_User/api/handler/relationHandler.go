package handler

import (
	"GoProject/model/vo"
	"GoProject/service/serviceImpl"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

// RelationAction
/*
	登录用户对其他用户进行关注或取关
*/
func RelationAction(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "ok",
	})
}

// FollowList
/*
	登录用户关注的所有用户列表
*/
func FollowList(ctx context.Context, c *app.RequestContext) {
	userId := c.Query("user_id")
	id, _ := strconv.ParseInt(userId, 10, 64)
	fsi := serviceImpl.FollowServiceImpl{}
	if UserInfoList, err := fsi.GetFollowListById(id); err == nil {
		c.JSON(consts.StatusOK, vo.FollowResponse{
			Response:     vo.Response{StatusCode: ResponseSuccess},
			UserInfoList: UserInfoList,
		})
	} else {
		c.JSON(consts.StatusOK, vo.FollowResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "Query UserInfo error"},
		})
	}
}

// FollowerList
/*
	所有关注登录的粉丝列表
*/
func FollowerList(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "ok",
	})
}

// FriendList
/*
	所有登录用户的好友列表
*/
func FriendList(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "ok",
	})
}
