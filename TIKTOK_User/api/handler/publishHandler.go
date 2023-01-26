package handler

import (
	"GoProject/model"
	"GoProject/model/vo"
	"GoProject/mw"
	"GoProject/service/serviceImpl"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"strconv"
)

// PublishAction
/*
	登录用户选择视频上传
*/
func PublishAction(ctx context.Context, c *app.RequestContext) {

	//user, _ := c.Get(mw.IdentityKey)
	videoTitle := c.PostForm("title")
	videoData, err := c.Request.FormFile("data")
	if err != nil {
		log.Print("can not get this filestream")
	}
	//videoTitle := c.Query("title")
	//reader := bytes.NewReader([]byte(videoData))
	psi := serviceImpl.PublishServiceImpl{}
	//user.(*model.User).Id
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

// PublishList
/*
	登录用户的视频发布列表
*/
func PublishList(ctx context.Context, c *app.RequestContext) {
	// 查询对象的userId
	userId := c.Query("user_id")
	// 通过token获取到的登录用户名
	user, _ := c.Get(mw.IdentityKey)
	id, _ := strconv.ParseInt(userId, 10, 64)

	psi := serviceImpl.PublishServiceImpl{}
	videoList, err := psi.GetVideoList(id, user.(*model.User).Id)
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
