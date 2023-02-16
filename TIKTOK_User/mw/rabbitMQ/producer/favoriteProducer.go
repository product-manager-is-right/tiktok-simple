package producer

import (
	"log"
	"strconv"
	"strings"
)

var favorProducer ExProducer

func InitFavorite() {
	initFavorProducer()
}

//	-> AddDelFavorite
//
// exchange :favor
//
//	-> UpdateFavoriteCnt
func initFavorProducer() {
	pe, err := NewWithExchange("favor", "fanout")
	if err != nil {
		log.Fatal("Favorite生产者初始化失败")
	}
	favorProducer = pe
}

// SendFavoriteMessage
// 发送结构为 userId-videoId-type
// actionType : 1 -> add ; 0 -> delete
func SendFavoriteMessage(userId int64, videoId int64, actionType int) error {
	//using rabbitMQ to store the info
	sb := strings.Builder{}
	//使用最高的36，压缩一下
	sb.WriteString(strconv.FormatInt(userId, 36))
	sb.WriteString("-")
	sb.WriteString(strconv.FormatInt(videoId, 36))
	sb.WriteString("-")
	sb.WriteString(strconv.Itoa(actionType))
	if err := favorProducer.Publish(sb.String()); err != nil {
		log.Print(err)
		return err
	}

	return nil
}
