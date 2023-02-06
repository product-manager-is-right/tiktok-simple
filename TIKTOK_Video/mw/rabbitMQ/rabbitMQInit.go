package rabbitMQ

import (
	"github.com/streadway/amqp"
	"log"
)

const MqUrl = "amqp://guest:guest@120.25.2.146:5672/"

type RabbitMQ struct {
	conn  *amqp.Connection
	mqUrl string
}

var Rmq *RabbitMQ

// InitRabbitMQ 初始化RabbitMQ
func InitRabbitMQ() {
	Rmq = &RabbitMQ{
		mqUrl: MqUrl,
	}
	dial, err := amqp.Dial(Rmq.mqUrl)
	if err != nil {
		log.Fatalf("%s:%s\n", err, "rabbitMQ连接失败")
	}
	Rmq.conn = dial

}

// 关闭mq通道和mq的连接。
func (r *RabbitMQ) destroy() {

	if err := r.conn.Close(); err != nil {
		log.Print("close MQ fail")
		return
	}
}
