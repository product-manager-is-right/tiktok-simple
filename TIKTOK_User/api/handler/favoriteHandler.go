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

// FavoriteAction
/*
	赞操作，登录用户对视频的点赞和取消点赞操作
*/
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "ok",
	})
}

// FavoriteList
/*
	登录用户的所有点赞视频
*/
func FavoriteList(ctx context.Context, c *app.RequestContext) {

	//url获取的用户id
	userId := c.Query("user_id")
	// 通过token获取到的登录用户名
	//user, _ := c.Get(mw.IdentityKey)
	id, _ := strconv.ParseInt(userId, 10, 64)

	fsi := serviceImpl.FavoriteServiceImpl{}
	if vl, err := fsi.GetFavoriteVideosListByUserId(id, 23); err == nil {
		c.JSON(consts.StatusOK, vo.FavoriteListResponse{
			Response:  vo.Response{StatusCode: ResponseFail, StatusMsg: "查询点赞列表成功"},
			VideoList: vl,
		})
	} else {
		c.JSON(consts.StatusOK, vo.FavoriteListResponse{
			Response:  vo.Response{StatusCode: ResponseFail, StatusMsg: "查询点赞列表失败"},
			VideoList: vl,
		})
	}

}
