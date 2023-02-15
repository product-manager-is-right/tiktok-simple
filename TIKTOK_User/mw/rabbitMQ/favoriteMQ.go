package rabbitMQ

import (
	"TIKTOK_User/dal/mysql"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"strings"
)

const favoriteQueueName = "favorite"

// 关注田添加或取消的消费方式。
func consumerFavorite(msgs <-chan amqp.Delivery) {
	var err error
	var userId, videoId int64
	var actionType int
	log.Println("开始消费")
	for d := range msgs {
		// 参数解析。
		params := strings.Split(fmt.Sprintf("%s", d.Body), "-")
		if len(params) != 3 {
			log.Println("favorite队列收到错误消息", err.Error())
			continue
		}
		if userId, err = strconv.ParseInt(params[0], 36, 64); err != nil {
			log.Println("favorite队列收到错误消息", err.Error())
			continue
		}
		if videoId, err = strconv.ParseInt(params[1], 36, 64); err != nil {
			log.Println("favorite队列收到错误消息", err.Error())
			continue
		}
		if actionType, err = strconv.Atoi(params[2]); err != nil {
			log.Println("favorite队列收到错误消息", err.Error())
			continue
		}
		log.Println("favorite接收点赞操作消息：用户", userId, "视频", videoId, "执行", actionType)
		switch actionType {
		case 0:
			if err = mysql.DeleteFavorite(userId, videoId); err != nil {
				log.Println("favorite队列消费者操作失败:", err.Error())
			}
		case 1:
			if _, err = mysql.CreateNewFavorite(userId, videoId); err != nil {
				log.Println("favorite队列消费者操作失败", err.Error())
			}
		}

	}
}

var RmqFavorite *MyMessageQueue

// InitFavoriteRabbitMQ 初始化rabbitMQ连接。
func InitFavoriteRabbitMQ() {
	RmqFavorite = NewRabbitMQSimple(favoriteQueueName)

	go RmqFavorite.ConsumeWithEx(consumerFavorite, "favor")

}
