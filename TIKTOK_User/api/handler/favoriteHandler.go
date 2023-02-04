package handler

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
	"TIKTOK_User/service"
	"TIKTOK_User/service/serviceImpl"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

// FavoriteAction
/*
	赞操作，登录用户对视频的点赞和取消点赞操作
*/
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	//url获取的用户id、视频id
	userId := c.Query("user_id")
	videoId := c.Query("video_id")
	// 通过token获取到的登录用户名
	//user, _ := c.Get(mw.IdentityKey)
	userid, _ := strconv.ParseInt(userId, 10, 64)
	videoid, _ := strconv.ParseInt(videoId, 10, 64)

	fsi := service.NewFavoriteServiceInstance()

	isFavorite, err := mysql.GetIsFavorite(userid, videoid)
	if err != nil {
		return
	}
	if isFavorite == false {
		fes, err := fsi.CreateNewFavorite(userid, videoid)
		if fes != -1 && err == nil {
			c.JSON(consts.StatusOK, vo.FavorVideoResponse{
				Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "点赞成功"},
			})
		} else {
			c.JSON(consts.StatusOK, vo.FavorVideoResponse{
				Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "点赞失败"},
			})
		}
	} else {
		err := fsi.DeleteFavorite(userid, videoid)
		if err == nil {
			//返回格式
			c.JSON(consts.StatusOK, vo.FavorVideoResponse{
				Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "取消点赞成功"},
			})
		} else {
			c.JSON(consts.StatusOK, vo.FavorVideoResponse{
				Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "取消点赞失败"},
			})
		}
	}

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

// IsFavorite
// 判断是否为喜欢接口,返回FavoriteInfoResponse对象
func IsFavorite(ctx context.Context, c *app.RequestContext) {
	userIdInfo := c.Query("userId")
	videoIdInfo := c.Query("videoId")
	if userIdInfo == "" || videoIdInfo == "" {
		c.JSON(consts.StatusOK, vo.FavoriteInfoResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "query userid or videoId empty"},
		})
		return
	}
	var userId, videoId int64
	var err error
	if userId, err = strconv.ParseInt(userIdInfo, 10, 64); err != nil {
		c.JSON(consts.StatusOK, vo.FavoriteInfoResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "query userid not a number"},
		})
		return
	}
	if videoId, err = strconv.ParseInt(videoIdInfo, 10, 64); err != nil {
		c.JSON(consts.StatusOK, vo.FavoriteInfoResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "query videoId not a number"},
		})
		return
	}
	fsi := service.NewFavoriteServiceInstance()
	var favorite bool
	if favorite, err = fsi.IsFavorite(userId, videoId); err != nil {
		c.JSON(consts.StatusOK, vo.FavoriteInfoResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "error happened"},
		})
		return
	}
	c.JSON(consts.StatusOK, vo.FavoriteInfoResponse{
		Response: vo.Response{
			StatusCode: ResponseSuccess,
			StatusMsg:  "query success",
		},
		IsFavorite: favorite,
	})

}
