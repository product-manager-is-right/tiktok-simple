package consumer

import (
	"TIKTOK_Video/dal/mysql"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"strings"
)

// InitFavorite 初始化FavoriteMq消费者。
func InitFavorite() {
	updateFavoriteCnt()
}

func updateFavoriteCnt() {
	c, err := NewWithExchange(consumerName, "favor", "fanout", "UpdateFavoriteCnt")
	if err != nil {
		log.Fatal("FavoriteConsumer消费者创建失败")
	}
	msg, err := c.Consume()
	if err != nil {
		log.Fatal("FavoriteConsumer消费失败")
	}

	go consumerFavor(msg)
}

// 点赞数+1 -1的消费方式。
func consumerFavor(msgs <-chan amqp.Delivery) {
	var err error
	var videoId int64
	var action int
	log.Println("updateFavoriteCnt : 开始消费")
	for d := range msgs {
		// 参数解析。
		params := strings.Split(fmt.Sprintf("%s", d.Body), "-")
		if len(params) != 3 {
			log.Println("favorite队列收到错误消息", err.Error())
			continue
		}
		if videoId, err = strconv.ParseInt(params[1], 36, 64); err != nil {
			log.Println("favorite队列收到错误消息", err.Error())
			continue
		}
		if action, err = strconv.Atoi(params[2]); err != nil {
			log.Println("favorite队列收到错误消息", err.Error())
			continue
		}
		log.Println("Favorite接收点赞操作消息：视频", videoId, "执行类型", action)
		switch action {
		case 1:
			if err = mysql.IncrementFavoriteCount(videoId); err != nil {
				log.Println("FavoriteCnt队列消费者+1失败:", err.Error())
			}
		case 0:
			if err = mysql.DecrementFavoriteCount(videoId); err != nil {
				log.Println("FavoriteCnt队列消费者-1失败", err.Error())
			}
		}

	}
}
