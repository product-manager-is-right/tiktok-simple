package handler

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
	"TIKTOK_User/service"
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
	userfromId := c.Query("user_id_from")
	usertoId := c.Query("user_id_to")
	//relationactType := c.Query("cancel")
	// 通过token获取到的登录用户名
	//user, _ := c.Get(mw.IdentityKey)
	userfromid, _ := strconv.ParseInt(userfromId, 10, 64)
	usertoid, _ := strconv.ParseInt(usertoId, 10, 64)
	//关注服务
	fsi := service.NewCommentServiceInstance()
	//关注方法
	isFollow, err := mysql.GetIsFollow(usertoid, userfromid)
	if err != nil {
		return
	}
	if isFollow == false {
		res, err := fsi.CreateNewRelation(userfromid, usertoid)
		if res != -1 && err == nil {
			//返回格式
			c.JSON(consts.StatusOK, vo.FollowActionResponse{
				Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "关注成功"},
			})
		} else {
			c.JSON(consts.StatusOK, vo.FollowActionResponse{
				Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "关注失败"},
			})
		}
	} else {
		err := fsi.DeleteRelation(userfromid, usertoid)
		if err == nil {
			//返回格式
			c.JSON(consts.StatusOK, vo.FollowActionResponse{
				Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "取关成功"},
			})
		} else {
			c.JSON(consts.StatusOK, vo.FollowActionResponse{
				Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "取关失败"},
			})
		}
	}

}

// FollowList
/*
	登录用户关注的所有用户列表
*/
func FollowList(ctx context.Context, c *app.RequestContext) {
	userId := c.Query("user_id")
	id, _ := strconv.ParseInt(userId, 10, 64)
	fsi := service.NewCommentServiceInstance()
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
	fsi := service.NewCommentServiceInstance2()
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
		c.JSON(http.StatusOK, vo.RelationResponse{
			Response:     vo.Response{StatusCode: ResponseFail, StatusMsg: "query user id empty"},
			UserInfoList: nil,
		})
		return
	}
	userId, err := strconv.ParseInt(u, 10, 64)
	// 用户id解析出错。
	if nil != err {
		c.JSON(http.StatusOK, vo.RelationResponse{
			Response:     vo.Response{StatusCode: ResponseSuccess, StatusMsg: "用户id格式错误"},
			UserInfoList: nil,
		})
		return
	}
	// 正常获取好友列表
	fsi := service.NewCommentServiceInstance()
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
