package mw

import (
	"TIKTOK_Video/mw/rabbitMQ/consumer"
	"TIKTOK_Video/mw/rabbitMQ/producer"
)

// InitRabbitMq Init 在此处创建生产者和消费者
func InitRabbitMq() {
	// 消费者
	{
		consumer.InitComment()
		consumer.InitFavorite()
	}

	// 生产者
	{
		producer.InitComment()
	}
}
