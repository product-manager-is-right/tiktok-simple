package rabbitMQ

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"strings"
)

// 关系添加的消费方式。
func consumerFollowAdd(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		// 参数解析。
		params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
		if len(params) != 2 {
			log.Println("follow_add队列收到错误消息")
			continue
		}
		userId, _ := strconv.Atoi(params[0])
		targetId, _ := strconv.Atoi(params[1])
		// 日志记录。
		log.Println("接受到添加关注消息：", userId, "关注", targetId)

		Follow := model.Follow{UserIdTo: int64(targetId), UserIdFrom: int64(userId)}

		if err := mysql.DB.Create(&Follow).Error; err != nil {
			log.Println(err)
		}
		/*
			sql := fmt.Sprintf("CALL CreateNewRelation(%v,%v)", targetId, userId)
			if err := mysql.DB.Raw(sql).Scan(nil).Error; nil != err {
			// 执行出错，打印日志。
			log.Println(err.Error())
			}
		*/
		// 执行SQL，注必须scan，该SQL才能被执行。

	}
}

// 关系删除的消费方式。
func consumerFollowDel(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		// 参数解析。
		params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
		if len(params) != 2 {
			log.Println("follow_del队列收到错误消息")
			continue
		}
		userId, _ := strconv.Atoi(params[0])
		targetId, _ := strconv.Atoi(params[1])
		// 日志记录。
		follow := &model.Follow{}
		log.Println("接受到取消关注消息：", userId, "关注", targetId)
		if err := mysql.DB.Model(follow).Where("user_id_to = ?", targetId).Where("user_id_from = ?", userId).
			Delete(follow).Error; err != nil {
			log.Fatal("更新失败")
		}
		/*
				sql := fmt.Sprintf("CALL UpdateRelation(%v,%v)", targetId, userId)
			log.Printf("消费队列执行删除关系。SQL如下：%s", sql)
			// 执行SQL，注必须scan，该SQL才能被执行。
			if err := mysql.DB.Raw(sql).Scan(nil).Error; nil != err {
				// 执行出错，打印日志。
				log.Println(err.Error())
			}
		*/

		// 再删Redis里的信息，防止脏数据，保证最终一致性。
		//updateRedisWithDel(userId, targetId)
	}
}

var RmqFollowAdd *MyMessageQueue
var RmqFollowDel *MyMessageQueue

// InitFollowRabbitMQ 初始化rabbitMQ连接。
func InitFollowRabbitMQ() {
	RmqFollowAdd = NewRabbitMQSimple("follow_add")
	//RmqFollowAdd = NewFollowRabbitMQ("follow_add")
	go RmqFollowAdd.ConsumeWithEx(consumerFollowAdd, "follow")

	RmqFollowDel = NewRabbitMQSimple("follow_del")
	//RmqFollowDel = NewFollowRabbitMQ("follow_del")
	go RmqFollowDel.ConsumeWithEx(consumerFollowDel, "follow")
}
