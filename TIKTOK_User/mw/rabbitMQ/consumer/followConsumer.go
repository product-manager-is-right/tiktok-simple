package consumer

import (
	"TIKTOK_User/dal/mysql"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"strings"
)

// InitFollow 初始化FollowMq消费者。
func InitFollow() {
	addDelFollow()
}

func addDelFollow() {
	c, err := NewWithSimple(consumerName, "AddDelFollow")
	msg, err := c.Consume()
	if err != nil {
		log.Fatal("FollowConsumer消费者创建失败")
	}
	go consumerAddDelFollow(msg)
}

// 关系添加的消费方式。
func consumerAddDelFollow(msg <-chan amqp.Delivery) {
	var err error
	var userIdFrom, videoIdTo int64
	var actionType int
	log.Println("addDelFollow : 开始消费")
	for d := range msg {
		// 参数解析。
		params := strings.Split(fmt.Sprintf("%s", d.Body), "-")
		if len(params) != 3 {
			log.Println("follow队列收到错误消息", err.Error())
			continue
		}
		if userIdFrom, err = strconv.ParseInt(params[0], 36, 64); err != nil {
			log.Println("follow队列收到错误消息", err.Error())
			continue
		}
		if videoIdTo, err = strconv.ParseInt(params[1], 36, 64); err != nil {
			log.Println("follow队列收到错误消息", err.Error())
			continue
		}
		if actionType, err = strconv.Atoi(params[2]); err != nil {
			log.Println("follow队列收到错误消息", err.Error())
			continue
		}
		log.Println("follow接收关注操作消息：from", userIdFrom, "to", videoIdTo, "执行", actionType)
		switch actionType {
		case 0:
			if err = mysql.DeleteRelation(videoIdTo, userIdFrom); err != nil {
				log.Println("favorite队列消费者操作失败:", err.Error())
			}
		case 1:
			if err = mysql.CreateNewRelation(videoIdTo, userIdFrom); err != nil {
				log.Println("favorite队列消费者操作失败", err.Error())
			}
		}

	}
}
