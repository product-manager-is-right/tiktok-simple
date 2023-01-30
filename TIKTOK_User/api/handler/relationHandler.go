package handler

import (
	"GoProject/model/vo"
	"GoProject/service/serviceImpl"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
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
	c.JSON(consts.StatusOK, utils.H{
		"message": "ok",
	})

	//// 查询对象的userId和其他用户的id
	//userId, err1 := strconv.ParseInt(c.GetString("user_id_from"), 10, 64)
	//toUserId, err2 := strconv.ParseInt(c.Query("user_id_to"), 10, 64)
	//actionType, err3 := strconv.ParseInt(c.Query("cancel"), 10, 64)
	//fmt.Println(userId, toUserId, actionType)
	//// 传入参数格式有问题。
	//if nil != err1 || nil != err2 || nil != err3 || actionType < 1 || actionType > 2 {
	//	fmt.Printf("fail")
	//	c.JSON(http.StatusOK, vo.FollowerResponse{
	//		Response: vo.Response{StatusCode: ResponseSuccess},
	//		//UserInfoList: "id type error",
	//	})
	//	return
}

// 正常处理，实例化对象
//	fsi := serviceImpl.GetFollowerListById(userId)
//	switch {
//	// 关注
//	case 1 == actionType:
//		go fsi.AddFollowRelation(userId, toUserId)
//	// 取关
//	case 2 == actionType:
//		go fsi.DeleteFollowRelation(userId, toUserId)
//	}
//	log.Println("关注、取关成功。")
//	c.JSON(http.StatusOK, RelationActionResp{
//		Response{
//			StatusCode: 0,
//			StatusMsg:  "OK",
//		},
//	})

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
