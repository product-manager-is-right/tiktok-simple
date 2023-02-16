package rabbitMQ

import (
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

var rabbitConn *amqp.Connection

var once = sync.Once{}

// GetConn InitRabbitMQ 创建RabbitMQ的连接，单例模式，只需要一个Connection。
func GetConn() *amqp.Connection {
	once.Do(func() {
		conn, err := amqp.Dial(viper.GetString("rabbitmq.url"))
		if err != nil {
			log.Fatal("rabbit Connection 失败")
		}
		rabbitConn = conn
	})
	return rabbitConn
}

// NewChannel 创建channel
func NewChannel() *amqp.Channel {
	ch, err := GetConn().Channel()
	if err != nil {
		log.Fatal("创建mq信道失败")
	}
	return ch
}

func CloseConn() {
	if rabbitConn != nil && !rabbitConn.IsClosed() {
		destroy()
	}
}

// Destroy 关闭mq通道和mq的连接。
func destroy() {
	if err := rabbitConn.Close(); err != nil {
		log.Print("close MQ fail")
		return
	}
}
