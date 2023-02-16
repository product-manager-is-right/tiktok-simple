package test

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/mw/rabbitMQ/consumer"
	"TIKTOK_User/mw/rabbitMQ/producer"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
)

func TestProducerWithEx(t *testing.T) {
	pe, _ := producer.NewWithExchange("test", "fanout")
	err := pe.Publish("haha")
	assert.Nil(t, err)
}

func TestConsumerWithEx(t *testing.T) {
	mysql.Init()
	ce, _ := consumer.NewWithExchange("user", "test", "fanout", "o")
	msg, _ := ce.Consume()
	ce2, _ := consumer.NewWithExchange("user", "test", "fanout", "oo")
	msg2, _ := ce2.Consume()
	var forever chan interface{}
	go func() {
		for v := range msg {
			println(string(v.Body))
		}
	}()
	for v := range msg2 {
		println(string(v.Body))
	}
	<-forever
}
