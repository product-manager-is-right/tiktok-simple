package consumer

import (
	"TIKTOK_User/dal/mysql"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"strings"
)

// InitFavorite 初始化FavoriteMq消费者。
func InitFavorite() {
	addDelFavorite()
}

func addDelFavorite() {
	c, err := NewWithExchange(consumerName, "favor", "fanout", "AddDelFavorite")
	if err != nil {
		log.Fatal("FavoriteConsumer消费者创建失败")
	}
	msg, err := c.Consume()
	if err != nil {
		log.Fatal("FavoriteConsumer消费失败")
	}

	go consumerFavor(msg)
}

// 关注田添加或取消的消费方式。
func consumerFavor(msg <-chan amqp.Delivery) {
	var err error
	var userId, videoId int64
	var actionType int
	log.Println("AddDelFavorite : 开始消费")
	for d := range msg {
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
