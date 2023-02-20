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

// MessageAction 发送消息api
func MessageAction(ctx context.Context, c *app.RequestContext) {
	ownerId, _ := c.Get(mw.IdentityKey)
	t := c.Query("to_user_id")
	toUserId, err := strconv.ParseInt(t, 10, 64)
	content := c.Query("content")
	if t == "" || content == "" || err != nil {
		c.JSON(consts.StatusOK, vo.MessageActionResponse{
			Response: vo.Response{
				StatusCode: ResponseFail,
				StatusMsg:  "请求参数不正确",
			},
		})
		return
	}

	msi := serviceImpl.MessageServiceImpl{}
	if err := msi.SendMessage(toUserId, ownerId.(int64), content); err != nil {
		c.JSON(consts.StatusOK, vo.MessageActionResponse{
			Response: vo.Response{
				StatusCode: ResponseFail,
				StatusMsg:  "发送消息错误：" + err.Error(),
			},
		})
		return
	}
	c.JSON(consts.StatusOK, vo.MessageActionResponse{
		Response: vo.Response{
			StatusCode: ResponseSuccess,
			StatusMsg:  "发送消息成功",
		},
	})
}

// MessageChat 聊天记录api
func MessageChat(ctx context.Context, c *app.RequestContext) {
	ownerId, _ := c.Get(mw.IdentityKey)
	t := c.Query("to_user_id")
	//TODO:处理pre_msg_time
	preT := c.Query("pre_msg_time")
	preTime, err := strconv.ParseInt(preT, 10, 64)
	toUserId, err := strconv.ParseInt(t, 10, 64)
	if t == "" || err != nil {
		c.JSON(consts.StatusOK, vo.MessageActionResponse{
			Response: vo.Response{
				StatusCode: ResponseFail,
				StatusMsg:  "请求参数不正确",
			},
		})
		return
	}
	msi := serviceImpl.MessageServiceImpl{}
	messageList, err := msi.GetMessage(toUserId, ownerId.(int64), preTime)
	if err != nil {
		c.JSON(consts.StatusOK, vo.ChatResponse{
			Response: vo.Response{
				StatusCode: ResponseFail,
				StatusMsg:  "获取聊天列表错误" + err.Error(),
			},
		})
		return
	}
	c.JSON(consts.StatusOK, vo.ChatResponse{
		Response:    vo.Response{StatusCode: ResponseSuccess, StatusMsg: "获取聊天列表成功"},
		MessageList: messageList,
	})
}
