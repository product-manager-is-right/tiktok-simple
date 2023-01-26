// Code generated by hertz generator.

package handler

import (
	"TIKTOK_Video/model/vo"
	"TIKTOK_Video/mw"
	"TIKTOK_Video/service/ServiceImpl"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/http"
	"strconv"
	"time"
)

const (
	ResponseSuccess = 0
	ResponseFail    = 1
)

/*
视频流接口
*/

func Feed(ctx context.Context, c *app.RequestContext) {
	var userName string
	var err error

	queryTime := c.Query("latest_time")
	token := c.Query("token")

	// 如果token不是空的，则可能登录状态，调用jwt鉴权。
	// 1. 鉴权失败，username为空，仍可以调用feed流
	// 2. 鉴权成功，可获取username
	// 如果token是空的，是未登录状态，username为空，仍可以调用feed流
	if token != "" {
		c.Next(ctx)
		if user, _ := c.Get(mw.IdentityKey); user != nil {
			userName = user.(string)
		}
	}
	c.Abort()

	var latestTime int64
	if queryTime != "" {
		latestTime, err = strconv.ParseInt(queryTime, 10, 64)
	} else {
		latestTime = time.Now().Unix()
	}

	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: ResponseFail,
			StatusMsg:  "时间戳请求错误",
		})
		return
	}
	vsi := ServiceImpl.VideoServiceImpl{}
	videoInfoList, nextTime, err := vsi.GetVideoInfosByLatestTime(latestTime, userName)
	if err != nil {
		c.JSON(http.StatusOK, vo.Response{
			StatusCode: ResponseFail,
			StatusMsg:  "获取视频流失败：" + err.Error(),
		})
		return
	}

	c.JSON(consts.StatusOK, vo.FeedResponse{
		Response:  vo.Response{StatusCode: ResponseSuccess},
		VideoList: videoInfoList,
		NextTime:  nextTime,
	})
}

/*
	评论操作接口
*/
