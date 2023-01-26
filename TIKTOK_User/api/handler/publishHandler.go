package handler

import (
	"GoProject/dal/mysql"
	"GoProject/model"
	"GoProject/model/vo"
	"GoProject/mw"
	"GoProject/service/serviceImpl"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

// PublishAction
/*
	登录用户选择视频上传
*/
func PublishAction(ctx context.Context, c *app.RequestContext) {
	user, _ := c.Get(mw.IdentityKey)
	videoData := c.Query("data")
	videoTitle := c.Query("title")

	psi := serviceImpl.PublishServiceImpl{}
	if err := psi.PublishVideo(user.(*model.User).Id, []byte(videoData), videoTitle); err != nil {
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

	// 通过token获取到的登录用户名，并通过sql查到userID
	userName, _ := c.Get(mw.IdentityKey)
	user, _ := mysql.GetUserByUserName(userName.(string))

	id, _ := strconv.ParseInt(userId, 10, 64)

	psi := serviceImpl.PublishServiceImpl{}
	videoList, err := psi.GetVideoList(id, user.Id)
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
