package rabbitMQ

import (
	"TIKTOK_Video/dal/mysql"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"strings"
)

const remoteFavorQueueName = "favorRemote"

// 关注田添加或取消的消费方式。
func consumerFavor(msgs <-chan amqp.Delivery) {
	var err error
	var videoId int64
	var action int
	for d := range msgs {
		// 参数解析。
		params := strings.Split(fmt.Sprintf("%s", d.Body), "-")
		if len(params) != 2 {
			log.Println("remoteFavorite队列收到错误消息", err.Error())
			continue
		}
		if videoId, err = strconv.ParseInt(params[1], 36, 64); err != nil {
			log.Println("remoteFavorite队列收到错误消息", err.Error())
			continue
		}
		if action, err = strconv.Atoi(params[2]); err != nil {
			log.Println("remoteFavorite队列收到错误消息", err.Error())
			continue
		}
		log.Println("remoteFavorite接收点赞操作消息：视频", videoId, "执行类型", action)
		switch action {
		case 0:
			if err = mysql.IncrementFavoriteCount(videoId); err != nil {
				log.Println("remoteFavorite队列消费者+1失败:", err.Error())
			}
		case 1:
			if err = mysql.DecrementFavoriteCount(videoId); err != nil {
				log.Println("favorite队列消费者操作-1失败", err.Error())
			}
		}

	}
}

var RmqFavorite *MyMessageQueue

// InitFavoriteRabbitMQ 初始化rabbitMQ连接。
func InitFavoriteRabbitMQ() {
	RmqFavorite = NewRabbitMQSimple(remoteFavorQueueName)

	go RmqFavorite.Consume(consumerFavor)

}
