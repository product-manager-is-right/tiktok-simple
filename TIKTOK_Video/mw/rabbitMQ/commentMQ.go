package rabbitMQ

import (
	"TIKTOK_Video/dal/mysql"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"strings"
)

var RmqComment *MyMessageQueue
var commentQueueName = "comment"

// InitCommentRabbitMQ 初始化删除队列的rabbitMQ连接。
func InitCommentRabbitMQ() {
	RmqComment = NewRabbitMQSimple(commentQueueName)

	RmqComment.Consume(consumerComment)

}

func consumerComment(msgs <-chan amqp.Delivery) {
	var err error
	var userId, videoId, commentId int64
	for d := range msgs {
		// 参数解析。
		params := strings.Split(fmt.Sprintf("%s", d.Body), "-")
		if len(params) != 3 {
			log.Println("favorite队列收到错误消息", err.Error())
			continue
		}
		if commentId, err = strconv.ParseInt(params[0], 36, 64); err != nil {
			log.Println("favorite队列收到错误消息", err.Error())
			continue
		}
		if userId, err = strconv.ParseInt(params[1], 36, 64); err != nil {
			log.Println("favorite队列收到错误消息", err.Error())
			continue
		}
		if videoId, err = strconv.ParseInt(params[2], 36, 64); err != nil {
			log.Println("favorite队列收到错误消息", err.Error())
			continue
		}

		log.Println("favorite接收点赞操作消息：用户", userId, "视频", videoId, "评论", commentId)
		//发送失败。自动同步操作数据库
		if err = mysql.DeleteCommentByCommentId(commentId, userId); err != nil {
			log.Println("comment队列消费者删除评论失败:", err.Error())
			continue
		}
		if err = mysql.DecrementCommentCount(videoId); err != nil {
			log.Println("comment队列消费者评论数减一失败:", err.Error())
		}
	}

}
