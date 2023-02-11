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

type FollowMQ struct {
	RabbitMQ
	channel   *amqp.Channel
	queueName string
	exchange  string
	key       string
}

// NewFollowRabbitMQ 获取followMQ的对应队列。
func NewFollowRabbitMQ(queueName string) *FollowMQ {
	followMQ := &FollowMQ{
		RabbitMQ:  *Rmq,
		queueName: queueName,
	}

	cha, err := followMQ.conn.Channel()
	followMQ.channel = cha
	if err != nil {
		log.Fatal(err, "获取通道失败")
	}

	return followMQ
}

// Publish 配置follower
func (f *FollowMQ) Publish(message string) error {
	//创建队列
	_, err := f.channel.QueueDeclare(
		f.queueName,
		//是否持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		return err
	}
	messages := []byte(message)
	err = f.channel.Publish(
		f.exchange,  //交换机名称
		f.queueName, //队列名称
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        messages,
		})
	if err != nil {
		log.Print("publish message failed")
		return err
	}
	return nil
}
func (f *FollowMQ) Consumer() {

	_, err := f.channel.QueueDeclare(f.queueName, false, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	//2、接收消息
	msgs, err := f.channel.Consume(
		f.queueName,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//消息队列是否阻塞
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	switch f.queueName {
	case "follow_add":
		go f.consumerFollowAdd(msgs)
	case "follow_del":
		go f.consumerFollowDel(msgs)

	}

	log.Printf("Waiting for messagees,To exit press CTRL+C")

	<-forever

}

// 关系添加的消费方式。
func (f *FollowMQ) consumerFollowAdd(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		// 参数解析。
		params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
		userId, _ := strconv.Atoi(params[0])
		targetId, _ := strconv.Atoi(params[1])
		// 日志记录。

		Follow := model.Follow{UserIdTo: int64(targetId), UserIdFrom: int64(userId), Cancel: 0}

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
func (f *FollowMQ) consumerFollowDel(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		// 参数解析。
		params := strings.Split(fmt.Sprintf("%s", d.Body), " ")
		userId, _ := strconv.Atoi(params[0])
		targetId, _ := strconv.Atoi(params[1])
		// 日志记录。
		if err := mysql.DB.Model(&model.Follow{}).Where("user_id_to = ?", targetId).Where("user_id_from = ?", userId).
			Update("cancel", 0).Error; err != nil {
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

var RmqFollowAdd *FollowMQ
var RmqFollowDel *FollowMQ

// InitFollowRabbitMQ 初始化rabbitMQ连接。
func InitFollowRabbitMQ() {
	RmqFollowAdd = NewFollowRabbitMQ("follow_add")
	go RmqFollowAdd.Consumer()

	RmqFollowDel = NewFollowRabbitMQ("follow_del")
	go RmqFollowDel.Consumer()
}
