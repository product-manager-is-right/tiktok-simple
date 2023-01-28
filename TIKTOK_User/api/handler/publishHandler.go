package handler

import (
	"TIKTOK_User/model/vo"
	"TIKTOK_User/mw"
	"TIKTOK_User/service/serviceImpl"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

// PublishList
/*
	登录用户的视频发布列表
*/
func PublishList(ctx context.Context, c *app.RequestContext) {
	// 查询对象的userId
	queryUserId := c.Query("user_id")
	if queryUserId == "" {
		c.JSON(consts.StatusOK, vo.VideoInfoResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "query user_id empty"},
		})
		return
	}

	// 通过token获取到的登录用户名，并通过sql查到userID
	userId, _ := c.Get(mw.IdentityKey)

	id, _ := strconv.ParseInt(queryUserId, 10, 64)

	psi := serviceImpl.PublishServiceImpl{}
	videoList, err := psi.GetVideoList(id, userId.(int64))
	if err != nil {
		c.JSON(consts.StatusOK, vo.Response{
			StatusCode: ResponseFail,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, vo.VideoInfoResponse{
		Response:  vo.Response{StatusCode: ResponseSuccess},
		VideoList: videoList,
	})
}
