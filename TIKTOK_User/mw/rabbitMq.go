package mw

import (
	"TIKTOK_User/mw/rabbitMQ/consumer"
	"TIKTOK_User/mw/rabbitMQ/producer"
)

// InitRabbitMq Init 在此处创建生产者和消费者
func InitRabbitMq() {
	// 消费者
	{
		consumer.InitFavorite()
		consumer.InitFollow()
	}

	// 生产者
	{
		producer.InitFavorite()
		producer.InitFollow()
	}
}
