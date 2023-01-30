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
	////url获取的对方用户id、视频id
	//userId := c.Query("to_user_id")
	////actionType := c.Query("action_type")
	//// 通过token获取到的登录用户名
	////user, _ := c.Get(mw.IdentityKey)
	//userid, _ := strconv.ParseInt(userId, 10, 64)
	////关注服务
	//fsi := serviceImpl.FollowServiceImpl{}
	////关注方法
	//res, err := fsi.CreateNewRelation(userid)
	//if res != -1 && err == nil {
	//	//返回格式
	//	c.JSON(consts.StatusOK, vo.FollowActionResponse{
	//		Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "关注成功"},
	//	})
	//} else {
	//	c.JSON(consts.StatusOK, vo.FollowActionResponse{
	//		Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "关注失败"},
	//	})
	//}
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
			Response:     vo.Response{StatusCode: ResponseSuccess, StatusMsg: "Query UserInfo success"},
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
	userId := c.Query("user_id")
	id, _ := strconv.ParseInt(userId, 10, 64)
	fsi := serviceImpl.FollowerServiceImpl{}
	if UserInfoList, err := fsi.GetFollowerListById(id); err == nil {
		c.JSON(consts.StatusOK, vo.FollowResponse{
			Response:     vo.Response{StatusCode: ResponseSuccess, StatusMsg: "Query UserInfo success"},
			UserInfoList: UserInfoList,
		})
	} else {
		c.JSON(consts.StatusOK, vo.FollowResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "Query UserInfo error"},
		})
	}
}

// FriendList
/*
	所有登录用户的好友列表
*/
func FriendList(ctx context.Context, c *app.RequestContext) {
	u := c.Query("user_id")
	if u == "" {
		c.JSON(http.StatusOK, vo.FollowerResponse{
			Response:     vo.Response{StatusCode: ResponseFail, StatusMsg: "query user id empty"},
			UserInfoList: nil,
		})
		return
	}
	userId, err := strconv.ParseInt(u, 10, 64)
	// 用户id解析出错。
	if nil != err {
		c.JSON(http.StatusOK, vo.FollowerResponse{
			Response:     vo.Response{StatusCode: ResponseSuccess, StatusMsg: "用户id格式错误"},
			UserInfoList: nil,
		})
		return
	}
	// 正常获取好友列表
	fsi := serviceImpl.FollowServiceImpl{}
	users, err := fsi.GetFollowListById(userId)
	// 获取关注列表时出错。
	if err != nil {
		c.JSON(http.StatusOK, vo.FollowerResponse{
			Response:     vo.Response{StatusCode: ResponseSuccess, StatusMsg: "获取好友列表时出错:" + err.Error()},
			UserInfoList: nil,
		})
		return
	}
	// 成功获取到粉丝列表。
	log.Println("获取粉丝列表成功。")
	c.JSON(http.StatusOK, vo.FollowerResponse{
		Response:     vo.Response{StatusCode: ResponseSuccess, StatusMsg: "Query UserInfo success"},
		UserInfoList: users,
	})
}
