package handler

import (
	"TIKTOK_Video/model/vo"
	"TIKTOK_Video/mw"
	"TIKTOK_Video/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
)

func CommentAction(ctx context.Context, c *app.RequestContext) {
	var actionTypeStr string
	videoIdStr := c.Query("video_id")
	//videoIdStr为空字符串说明请求中没有带这个参数
	if videoIdStr == "" {
		returnEmptyResponse("need param named 'video_id'", c)
		return
	}
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		returnEmptyResponse("video_id error", c)
		return
	}

	userId, _ := c.Get(mw.IdentityKey)

	instance := service.NewCommentServiceInstance()

	actionTypeStr = c.Query("action_type")
	if actionTypeStr == "1" {
		commentText := c.Query("comment_text")
		if commentText == "" {
			returnEmptyResponse("need param named 'comment_text'", c)
			return
		}
		commentInfo, err := instance.InsertComment(commentText, videoId, userId.(int64))
		if err != nil {
			returnEmptyResponse(err.Error(), c)
			return
		}
		c.JSON(consts.StatusOK, vo.CommentResponse{
			Response: vo.Response{
				StatusCode: 0,
				StatusMsg:  "insert success",
			},
			Comment: *commentInfo,
		})
		return
	} else if actionTypeStr == "2" {
		commentIdStr := c.Query("comment_id")
		if commentIdStr == "" {
			returnEmptyResponse("need param named 'comment_id'", c)
			return
		}
		var commentId int64
		if commentId, err = strconv.ParseInt(commentIdStr, 10, 64); err != nil {
			returnEmptyResponse("comment_id error", c)
			return
		}
		if err = instance.DeleteCommentByCommentId(commentId, userId.(int64), videoId); err != nil {
			returnEmptyResponse(err.Error(), c)
		} else {
			c.JSON(consts.StatusOK, vo.Response{
				StatusCode: 0,
				StatusMsg:  "delete success",
			})
		}
		return
	}
	returnEmptyResponse("need param named 'action_type'", c)
	return
}

/*
CommentList
评论列表接口
*/
func CommentList(ctx context.Context, c *app.RequestContext) {
	//解析videoId
	videoIdStr := c.Query("video_id")
	//videoIdStr为空字符串说明请求中没有带这个参数
	if videoIdStr == "" {
		c.JSON(consts.StatusOK, vo.CommentListResponse{
			Response: vo.Response{
				StatusCode: -1,
				StatusMsg:  "need param named 'video_id'",
			},
			CommentList: nil,
		})
		return
	}

	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		c.JSON(consts.StatusOK, vo.CommentListResponse{
			Response: vo.Response{
				StatusCode: -1,
				StatusMsg:  "video_id error",
			},
			CommentList: nil,
		})
		return
	}

	userId, _ := c.Get(mw.IdentityKey)

	instance := service.NewCommentServiceInstance()
	commentInfos, err := instance.GetCommentListByVideoId(videoId, userId.(int64))
	if err != nil {
		c.JSON(consts.StatusOK, vo.CommentListResponse{
			Response: vo.Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
			CommentList: nil,
		})
		return
	}
	c.JSON(consts.StatusOK, vo.CommentListResponse{
		Response: vo.Response{
			StatusCode: 0,
			StatusMsg:  "query success",
		},
		CommentList: commentInfos,
	})

}

func returnEmptyResponse(msg string, c *app.RequestContext) {
	c.JSON(consts.StatusOK, vo.CommentListResponse{
		Response: vo.Response{
			StatusCode: -1,
			StatusMsg:  msg,
		},
		CommentList: nil,
	})
}
