package handler

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
	"TIKTOK_User/service"
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
	//url获取的对方用户id、视频id
	userFromId := c.Query("user_id_from")
	userToId := c.Query("user_id_to")
	//relationactType := c.Query("cancel")
	// 通过token获取到的登录用户名
	//user, _ := c.Get(mw.IdentityKey)
	userFrom, _ := strconv.ParseInt(userFromId, 10, 64)
	userTo, _ := strconv.ParseInt(userToId, 10, 64)
	//关注服务
	fsi := service.NewFollowServiceInstance()
	//关注方法
	isFollow, err := mysql.GetIsFollow(userTo, userFrom)
	if err != nil {
		log.Print("关注失败")
		c.JSON(consts.StatusOK, vo.FollowActionResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "关注失败"},
		})
	}
	if isFollow == false {
		res, err := fsi.CreateNewRelation(userFrom, userTo)
		if res == -1 && err != nil {
			//返回格式
			c.JSON(consts.StatusOK, vo.FollowActionResponse{
				Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "关注失败"},
			})
		}
		c.JSON(consts.StatusOK, vo.FollowActionResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "关注成功"},
		})
	} else {
		err := fsi.DeleteRelation(userFrom, userTo)
		if err != nil {
			c.JSON(consts.StatusOK, vo.FollowActionResponse{
				Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "取关失败"},
			})
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
	id, _ := strconv.ParseInt(userId, 10, 64)
	fsi := service.NewFollowServiceInstance()
	UserInfoList, err := fsi.GetFollowListById(id)
	if err != nil {
		c.JSON(consts.StatusOK, vo.FollowResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "Query UserInfo error"},
		})
	}
	c.JSON(consts.StatusOK, vo.FollowResponse{
		Response:     vo.Response{StatusCode: ResponseSuccess, StatusMsg: "Query UserInfo success"},
		UserInfoList: UserInfoList,
	})
}

// FollowerList
/*
	所有关注登录的粉丝列表
*/
func FollowerList(ctx context.Context, c *app.RequestContext) {
	userId := c.Query("user_id")
	id, _ := strconv.ParseInt(userId, 10, 64)
	fsi := &serviceImpl.FollowerServiceImpl{}
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
	userIdVar := c.Query("user_id")
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
	fsi := service.NewFollowServiceInstance()
	users, err := fsi.GetFollowListById(userId)
	// 获取关注列表时出错。
	if err != nil {
		c.JSON(http.StatusOK, vo.RelationResponse{
			Response:     vo.Response{StatusCode: ResponseSuccess, StatusMsg: "获取好友列表时出错:" + err.Error()},
			UserInfoList: nil,
		})
		return
	}
	// 成功获取到粉丝列表。
	log.Println("获取粉丝列表成功。")
	c.JSON(http.StatusOK, vo.RelationResponse{
		Response:     vo.Response{StatusCode: ResponseSuccess, StatusMsg: "Query UserInfo success"},
		UserInfoList: users,
	})
}
