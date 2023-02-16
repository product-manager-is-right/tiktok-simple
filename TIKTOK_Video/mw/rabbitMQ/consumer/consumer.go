package consumer

import (
	"TIKTOK_Video/mw/rabbitMQ"
	"github.com/streadway/amqp"
	"log"
)

const consumerName = "tiktok.video"

type Consumer struct {
	ch           *amqp.Channel
	consumerName string
	queueName    string
}

// NewWithSimple  创建简单模式消费者
func NewWithSimple(consumerName, queueName string) (Consumer, error) {
	ch := rabbitMQ.NewChannel()

	q, err := ch.QueueDeclare(
		queueName,
		//是否持久化
		true,
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
		return Consumer{}, err
	}
	return Consumer{ch, consumerName, q.Name}, nil
}

func (c *Consumer) Consume() (<-chan amqp.Delivery, error) {
	delivery, err := c.ch.Consume(
		c.queueName,
		//用来区分多个消费者
		c.consumerName,
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
		log.Println(err)
	}
	return delivery, err
}

type ExConsumer struct {
	ch           *amqp.Channel
	consumerName string
	queueName    string
}

// NewWithExchange  创建发布-订阅模式消费者
func NewWithExchange(consumerName, exchange, kind, queueName string) (ExConsumer, error) {
	ch := rabbitMQ.NewChannel()
	err := ch.ExchangeDeclare(
		exchange, // name
		kind,     // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Println("消费者channel创建错误:", err)
		return ExConsumer{}, err
	}
	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Println("消费者channel创建错误:", err)
		return ExConsumer{}, err
	}
	err = ch.QueueBind(
		q.Name,   // queue name
		"",       // routing key
		exchange, // exchange
		false,
		nil,
	)
	if err != nil {
		log.Println("消费者channel创建错误:", err)
		return ExConsumer{}, err
	}
	return ExConsumer{ch, consumerName, q.Name}, nil
}

func (c *ExConsumer) Consume() (<-chan amqp.Delivery, error) {
	msg, err := c.ch.Consume(
		c.queueName,    // queue
		c.consumerName, // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		log.Println(err)
	}
	return msg, err
}
