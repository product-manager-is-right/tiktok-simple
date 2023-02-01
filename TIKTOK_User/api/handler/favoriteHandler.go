package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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
	c.JSON(consts.StatusOK, utils.H{
		"message": "ok",
	})
}
func IsFavorite(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "ok",
	})
}
