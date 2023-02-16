package producer

import (
	"log"
	"strconv"
	"strings"
)

var delCommentProducer Producer

func InitComment() {
	initDelCommentProducer()
}

func initDelCommentProducer() {
	p, err := NewWithSimple("DelComment")
	if err != nil {
		log.Fatal("Comment生产者初始化失败")
	}
	delCommentProducer = p
}

// SendDelCommentMessage 发送顺序为commentId, userId, videoId
func SendDelCommentMessage(commentId, videoId int64) error {
	//使用最高的36，压缩一下
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatInt(commentId, 36))
	sb.WriteString("-")
	sb.WriteString(strconv.FormatInt(videoId, 36))
	if err := delCommentProducer.Publish(sb.String()); err != nil {
		log.Print(err)
		return err
	}
	return nil
}
