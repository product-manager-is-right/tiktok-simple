package producer

import (
	"TIKTOK_Video/mw/rabbitMQ"
	"github.com/streadway/amqp"
	"log"
)

type Producer struct {
	ch        *amqp.Channel
	queueName string
}

// NewWithSimple 创建简单模式生产者
func NewWithSimple(queueName string) (Producer, error) {
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
		return Producer{}, err
	}
	return Producer{ch, q.Name}, nil
}

func (p *Producer) Publish(message string) error {
	b := []byte(message)
	err := p.ch.Publish(
		"",
		p.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		})
	if err != nil {
		log.Print("publish message failed")
		return err
	}
	return nil
}

type ExProducer struct {
	ch       *amqp.Channel
	exchange string
}

// NewWithExchange 创建发布-订阅模式生产者
func NewWithExchange(exchange, kind string) (ExProducer, error) {
	ch := rabbitMQ.NewChannel()
	//申请交换机
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
		return ExProducer{}, err
	}

	return ExProducer{ch, exchange}, nil
}

func (p *ExProducer) Publish(message string) error {
	b := []byte(message)

	err := p.ch.Publish(
		p.exchange, //交换机名称
		"",         //routing key
		false,      //如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,      //如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		})
	if err != nil {
		log.Print("publish message failed")
		return err
	}
	return nil
}
