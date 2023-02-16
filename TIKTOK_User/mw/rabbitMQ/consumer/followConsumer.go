package consumer

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/mw/redis"
	"context"
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
	var userIdFrom, userIdTo int64
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
		if userIdTo, err = strconv.ParseInt(params[1], 36, 64); err != nil {
			log.Println("follow队列收到错误消息", err.Error())
			continue
		}
		if actionType, err = strconv.Atoi(params[2]); err != nil {
			log.Println("follow队列收到错误消息", err.Error())
			continue
		}
		log.Println("follow接收关注操作消息：from", userIdFrom, "to", userIdTo, "执行", actionType)
		switch actionType {
		case 0:
			if err = mysql.DeleteRelation(userIdTo, userIdFrom); err != nil {
				log.Println("favorite队列消费者操作失败:", err.Error())
			}
			// follow数据库已经改变，关注列表和粉丝列表都要删除对应的key， 重试机制保证删除
			strUserFromId := strconv.FormatInt(userIdFrom, 10)
			strUserToId := strconv.FormatInt(userIdTo, 10)
			for i := 0; i < redis.RetryTime; i++ {
				if _, err := redis.FollowList.Del(context.Background(), strUserFromId).Result(); err == nil {
					break
				}
			}
			for i := 0; i < redis.RetryTime; i++ {
				if _, err := redis.FollowerList.Del(context.Background(), strUserToId).Result(); err == nil {
					break
				}
			}
		case 1:
			if err = mysql.CreateRelation(userIdTo, userIdFrom); err != nil {
				log.Println("favorite队列消费者操作失败", err.Error())
			}
			// follow数据库已经改变，关注列表和粉丝列表都要删除对应的key， 重试机制保证删除
			strUserFromId := strconv.FormatInt(userIdFrom, 10)
			strUserToId := strconv.FormatInt(userIdTo, 10)
			for i := 0; i < redis.RetryTime; i++ {
				if _, err := redis.FollowList.Del(context.Background(), strUserFromId).Result(); err == nil {
					break
				}
			}
			for i := 0; i < redis.RetryTime; i++ {
				if _, err := redis.FollowerList.Del(context.Background(), strUserToId).Result(); err == nil {
					break
				}
			}
		}
	}
}
