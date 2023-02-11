package test

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/mw/rabbitMQ"
	"TIKTOK_User/service/serviceImpl"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"log"
	"testing"
)

func Test(t *testing.T) {
	// 初始化数据库
	mysql.Init()

	rabbitMQ.InitRabbitMQ()
	// 初始化Follow的相关消息队列，并开启消费。
	rabbitMQ.InitFollowRabbitMQ()
	// 初始化Like的相关消息队列，并开启消费。
	fsi := serviceImpl.FollowServiceImpl{}
	err := fsi.CreateNewRelation(29, 29)
	if err != nil {
		log.Print("there are some errors")
	}
	assert.Nil(t, err)

}
