package consumer

import (
	"TIKTOK_Video/dal/mysql"
	"TIKTOK_Video/mw/redis"
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"strings"
)

func InitComment() {
	delComment()
}

func delComment() {
	c, err := NewWithSimple(consumerName, "DelComment")
	if err != nil {
		log.Fatal("Comment消费者创建失败")
	}
	msg, err := c.Consume()
	if err != nil {
		log.Fatal("Comment接收消息失败")
	}

	go consumerDelComment(msg)
}

func consumerDelComment(msgs <-chan amqp.Delivery) {
	var err error
	var videoId, commentId int64
	log.Println("delComment : 开始消费")
	for d := range msgs {
		// 参数解析。
		params := strings.Split(fmt.Sprintf("%s", d.Body), "-")
		if len(params) != 2 {
			log.Println("comment队列收到错误消息", err.Error())
			continue
		}
		if commentId, err = strconv.ParseInt(params[0], 36, 64); err != nil {
			log.Println("comment队列收到错误消息", err.Error())
			continue
		}
		if videoId, err = strconv.ParseInt(params[1], 36, 64); err != nil {
			log.Println("comment队列收到错误消息", err.Error())
			continue
		}

		log.Println("comment接受消息:", "视频", videoId, "评论", commentId)
		// 这里应该是一个事务
		if err = mysql.DeleteCommentByCommentId(commentId); err != nil {
			log.Println("comment队列消费者删除评论失败:", err.Error())
		}
		if err = mysql.DecrementCommentCount(videoId); err != nil {
			log.Println("comment队列消费者评论数减一失败:", err.Error())
		}
		strVideoId := strconv.FormatInt(videoId, 10)
		for i := 0; i < redis.RetryTime; i++ {
			if _, err := redis.CommentList.Del(context.Background(), strVideoId).Result(); err == nil {
				break
			}
		}
		continue
	}

}
