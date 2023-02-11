package handler

import (
	"TIKTOK_User/model/vo"
	"TIKTOK_User/mw"
	"TIKTOK_User/service/serviceImpl"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	_ "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"net/http"
	"strconv"
)

// RelationAction
/*
	登录用户对其他用户进行关注或取关
*/
func RelationAction(ctx context.Context, c *app.RequestContext) {
	userTo := c.Query("to_user_id")
	actionType := c.Query("action_type")

	userFromId, _ := c.Get(mw.IdentityKey)
	userToId, _ := strconv.ParseInt(userTo, 10, 64)

	//关注服务
	fsi := serviceImpl.FollowServiceImpl{}
	if actionType == "1" {
		//err := fsi.CreateNewRelation(24, userToId)
		err := fsi.CreateNewRelation(userFromId.(int64), userToId)
		if err != nil {
			//返回格式
			c.JSON(consts.StatusOK, vo.FollowActionResponse{
				Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "关注失败:" + err.Error()},
			})
			return
		}
		c.JSON(consts.StatusOK, vo.FollowActionResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "关注成功"},
		})
	} else {
		err := fsi.DeleteRelation(userFromId.(int64), userToId)
		if err != nil {
			c.JSON(consts.StatusOK, vo.FollowActionResponse{
				Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "取关失败:" + err.Error()},
			})
			return
		}
		//返回格式
		c.JSON(consts.StatusOK, vo.FollowActionResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "取关成功"},
		})
	}

}

// FollowList
/*
	登录用户关注的所有用户列表
*/
func FollowList(ctx context.Context, c *app.RequestContext) {
	userId := c.Query("user_id")
	ownerId, _ := c.Get(mw.IdentityKey)
	id, _ := strconv.ParseInt(userId, 10, 64)

	fsi := serviceImpl.FollowServiceImpl{}
	UserInfoList, err := fsi.GetFollowListById(id, ownerId.(int64))
	if err != nil {
		c.JSON(consts.StatusOK, vo.FollowResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "获取粉丝列表失败:" + err.Error()},
		})
	}
	c.JSON(consts.StatusOK, vo.FollowResponse{
		Response:     vo.Response{StatusCode: ResponseSuccess, StatusMsg: "获取粉丝列表成功"},
		UserInfoList: UserInfoList,
	})
}

// FollowerList
/*
	所有关注登录的粉丝列表
*/
func FollowerList(ctx context.Context, c *app.RequestContext) {
	userId := c.Query("user_id")
	ownerId, _ := c.Get(mw.IdentityKey)
	id, _ := strconv.ParseInt(userId, 10, 64)

	fsi := &serviceImpl.FollowerServiceImpl{}
	if UserInfoList, err := fsi.GetFollowerListById(id, ownerId.(int64)); err == nil {
		c.JSON(consts.StatusOK, vo.FollowResponse{
			Response:     vo.Response{StatusCode: ResponseSuccess, StatusMsg: "获取粉丝列表成功"},
			UserInfoList: UserInfoList,
		})
	} else {
		c.JSON(consts.StatusOK, vo.FollowResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "获取粉丝列表失败：" + err.Error()},
		})
	}
}

// FriendList
/*
	所有登录用户的好友列表
*/
func FriendList(ctx context.Context, c *app.RequestContext) {
	userIdVar := c.Query("user_id")
	ownerId, _ := c.Get(mw.IdentityKey)

	if userIdVar == "" {
		c.JSON(http.StatusOK, vo.RelationResponse{
			Response:     vo.Response{StatusCode: ResponseFail, StatusMsg: "query user id empty"},
			UserInfoList: nil,
		})
		log.Fatal("get useId failed ")
	}
	userId, err := strconv.ParseInt(userIdVar, 10, 64)
	// 用户id解析出错。
	if nil != err {
		c.JSON(http.StatusOK, vo.RelationResponse{
			Response:     vo.Response{StatusCode: ResponseSuccess, StatusMsg: "用户id格式错误"},
			UserInfoList: nil,
		})
		return
	}
	// 正常获取好友列表
	fsi := serviceImpl.FriendServiceImpl{}
	users, err := fsi.GetFriendListById(userId, ownerId.(int64))

	if err != nil {
		c.JSON(http.StatusOK, vo.RelationResponse{
			Response:     vo.Response{StatusCode: ResponseSuccess, StatusMsg: "获取好友列表时出错:" + err.Error()},
			UserInfoList: nil,
		})
		return
	}

	c.JSON(http.StatusOK, vo.RelationResponse{
		Response:     vo.Response{StatusCode: ResponseSuccess, StatusMsg: "获取好友列表成功"},
		UserInfoList: users,
	})
}
