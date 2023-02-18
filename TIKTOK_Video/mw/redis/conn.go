package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"math/rand"
	"time"
)

const RetryTime = 3

var ctx = context.Background()
var CommentList *redis.Client

// Init 初始化Redis连接。
func Init() {
	CommentList = redis.NewClient(&redis.Options{
		Addr:     "120.25.2.146:6379",
		Password: "roots",
		DB:       3, // 评论列表信息存入 DB3.
	})
	_, err := CommentList.Ping(ctx).Result()
	if err != nil {
		log.Fatal("连接redis失败")
	}
	log.Println("redis初始化成功")
}

func SetExpiredTime() time.Duration {
	n := rand.Intn(30)
	n += 30
	return time.Duration(n) * time.Second
}
