package producer

import (
	"log"
	"strconv"
	"strings"
)

var AddDelFollowProducer Producer

func InitFollow() {
	initAddDelFollowProducer()
}

func initAddDelFollowProducer() {
	p, err := NewWithSimple("AddDelFollow")
	if err != nil {
		log.Fatal("Follow生产者初始化失败")
	}
	AddDelFollowProducer = p
}
func SendFollowMessage(userToId int64, userFromId int64, actionType int) error {
	//using rabbitMQ to store the info
	sb := strings.Builder{}
	//使用最高的36，压缩一下
	sb.WriteString(strconv.FormatInt(userFromId, 36))
	sb.WriteString("-")
	sb.WriteString(strconv.FormatInt(userToId, 36))
	sb.WriteString("-")
	sb.WriteString(strconv.Itoa(actionType))
	if err := AddDelFollowProducer.Publish(sb.String()); err != nil {
		log.Print(err)
		return err
	}

	return nil
}
