package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var ctx = context.Background()
var FollowList *redis.Client
var FollowerList *redis.Client

var FavoriteList *redis.Client // Set key:userId,value:VideoId

// Init 初始化Redis连接。
func Init() {
	FollowList = redis.NewClient(&redis.Options{
		Addr:     "120.25.2.146:6379",
		Password: "roots",
		DB:       0, // 关注列表信息存入 DB0.
	})
	_, err := FollowList.Ping(ctx).Result()
	if err != nil {
		log.Fatal("连接redis失败")
	}

	FollowerList = redis.NewClient(&redis.Options{
		Addr:     "120.25.2.146:6379",
		Password: "roots",
		DB:       1, // 粉丝列表信息存入 DB1.
	})
	_, err = FollowerList.Ping(ctx).Result()
	if err != nil {
		log.Fatal("连接redis失败")
	}

	FavoriteList = redis.NewClient(&redis.Options{
		Addr:     "120.25.2.146:6379",
		Password: "roots",
		DB:       2, // 点赞列表信息存入 DB2.
	})
	_, err = FavoriteList.Ping(ctx).Result()
	if err != nil {
		log.Fatal("连接redis失败")
	}

	log.Println("redis初始化成功")
}
