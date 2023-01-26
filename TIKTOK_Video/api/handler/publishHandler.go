package handler

import (
	"TIKTOK_Video/model/vo"
	"TIKTOK_Video/service/ServiceImpl"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
)

// PublishAction
/*
	登录用户选择视频上传
*/
func PublishAction(ctx context.Context, c *app.RequestContext) {
	// get the basic info from meta
	//user, _ := c.Get(mw.IdentityKey)
	videoTitle := c.PostForm("title")
	videoData, err := c.Request.FormFile("data")
	if err != nil {
		log.Print("can not get this filestream")
	}
	psi := ServiceImpl.PublishServiceImpl{}
	if err := psi.PublishVideo(24, videoData, videoTitle); err != nil {
		c.JSON(consts.StatusOK, vo.Response{
			StatusCode: ResponseFail,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(consts.StatusOK, vo.Response{
		StatusCode: ResponseSuccess,
		StatusMsg:  "publish success!",
	})

}
