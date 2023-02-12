package test

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/mw/rabbitMQ"
	"TIKTOK_User/service/serviceImpl"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"log"
	"testing"
	"time"
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
	//ageng: 我发现测试案例中mq好像还没有消费就断开了，适当睡眠一下观察终端的打印日志，不然会直接停掉所有协程
	time.Sleep(time.Second * 2)
	assert.Nil(t, err)

}
