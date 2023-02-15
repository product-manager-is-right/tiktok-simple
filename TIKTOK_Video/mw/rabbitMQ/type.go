package rabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//定义数据结构和通用的方法等

//有普适性，我就改了名字了

type MyMessageQueue struct {
	RabbitMQ
	channel *amqp.Channel
	//队列名称
	queueName string
	//交换机名称
	exchange string
	//bind Key 名称
	key string
}

// NewMyMessageQueue 创建MyMessageQueue结构体实例
func NewMyMessageQueue(queueName string, exchange string, key string) *MyMessageQueue {
	return &MyMessageQueue{RabbitMQ: *Rmq, queueName: queueName, exchange: exchange, key: key}
}

// NewRabbitMQSimple 创建简单模式下MyMessageQueue结构体实例
func NewRabbitMQSimple(queueName string) *MyMessageQueue {
	//创建RabbitMQ实例
	rabbitmq := NewMyMessageQueue(queueName, "", "")
	var err error

	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()

	if err != nil {
		log.Fatal(err, "获取通道失败")
	}

	return rabbitmq
}

// Publish 直接模式队列生产
func (r *MyMessageQueue) Publish(message string) error {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
		r.queueName,
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
	//调用channel 发送消息到队列中
	messages := []byte(message)
	err = r.channel.Publish(
		r.exchange,  //交换机名称
		r.queueName, //队列名称充当路由key
		false,       //如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,       //如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
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

// Consume simple 模式下消费者
func (r *MyMessageQueue) Consume(consumeMethod func(msgs <-chan amqp.Delivery)) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
		r.queueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//2. 接收消息
	msgs, err := r.channel.Consume(
		r.queueName,
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
		fmt.Println(err)
	}

	//不需要这个，这个是使用案例中为了让程序不停止的做法而已，web项目不需要
	//forever := make(chan bool)
	//启用协程处理消息
	go consumeMethod(msgs)

	log.Printf("Waiting for messagees,To exit press CTRL+C")
	//<-forever

}
func (r *MyMessageQueue) ConsumeWithEx(consumeMethod func(msgs <-chan amqp.Delivery)) {
	//申请交换机
	err := r.channel.ExchangeDeclare(
		"favor",  // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err = r.channel.QueueDeclare(
		r.queueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	err = r.channel.QueueBind(
		r.queueName, // queue name
		"",          // routing key
		"favor",     // exchange
		false,
		nil,
	)
	//2. 接收消息
	msgs, err := r.channel.Consume(
		r.queueName,
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
		fmt.Println(err)
	}

	//不需要这个，这个是使用案例中为了让程序不停止的做法而已，web项目不需要
	//forever := make(chan bool)
	//启用协程处理消息
	go consumeMethod(msgs)

	log.Printf("Waiting for messagees,To exit press CTRL+C")
	//<-forever

}
