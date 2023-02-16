package test

import (
	"TIKTOK_Video/configs"
	"TIKTOK_Video/dal/mysql"
	"TIKTOK_Video/mw/rabbitMQ"
	"TIKTOK_Video/service"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
	"time"
)

func TestCommentMQ(t *testing.T) {
	configs.ReadConfig(configs.TEST)
	mysql.Init()
	rabbitMQ.InitRabbitMQ()
	rabbitMQ.InitCommentRabbitMQ()
	var err error
	var tests = []struct {
		commentId int64
		userId    int64
		videoId   int64
	}{
		{55, 24, 3},
	}
	csi := service.NewCommentServiceInstance()
	for _, test := range tests {
		err = csi.DeleteCommentByCommentId(test.commentId, test.userId, test.videoId)
		assert.Nil(t, err)
	}
	time.Sleep(time.Second * 3)

}
