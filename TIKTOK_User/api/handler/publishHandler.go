package handler

import (
	"TIKTOK_User/model/vo"
	"TIKTOK_User/mw"
	"TIKTOK_User/service/serviceImpl"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
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

// PublishVideo /*
// 该接口为远端调用接口，将userId和videoId存储进ums数据库
func PublishVideo(ctx context.Context, c *app.RequestContext) {
	userIdInfo := c.Query("userId")
	videoIdInfo := c.Query("videoId")
	if userIdInfo == "" || videoIdInfo == "" {
		c.JSON(consts.StatusOK, vo.RegisterResponse{
			Response: vo.Response{StatusCode: ResponseFail, StatusMsg: "query userid or videoId empty"},
		})
	}

	userId, err := strconv.ParseInt(userIdInfo, 10, 64)
	if err != nil {
		log.Print("can not change userId into int64")
	}
	videoId, err := strconv.ParseInt(videoIdInfo, 10, 64)
	if err != nil {
		log.Print("can not change videoId into int64")
	}
	vsi := serviceImpl.PublishServiceImpl{}
	if err := vsi.PublishVideoInfo(userId, videoId); err != nil {
		c.JSON(consts.StatusOK, vo.Response{
			StatusCode: ResponseFail,
			StatusMsg:  "error :" + err.Error(),
		})
		return
	}
	c.JSON(consts.StatusOK, vo.Response{
		StatusCode: ResponseSuccess,
		StatusMsg:  "upload info successful",
	})
	/*
		usi := serviceImpl.UserServiceImpl{}
		if _, err := usi.CreateUserByNameAndPassword(username, password); err != nil {
			c.JSON(consts.StatusOK, vo.RegisterResponse{
				Response: vo.Response{
					StatusCode: ResponseFail,
					StatusMsg:  "error :" + err.Error()},
			})
			c.Abort()
		}
	*/

}
