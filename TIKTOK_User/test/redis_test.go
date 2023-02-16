package test

import (
	"TIKTOK_User/mw"
	"TIKTOK_User/mw/redis"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"testing"
)

func TestRedisAdd(t *testing.T) {
	mw.RedisInit()
	_, err := redis.FavoriteList.SAdd(context.Background(), "1", "1", "2", "3").Result()
	assert.Nil(t, err)
}

func TestRedisDel(t *testing.T) {
	mw.RedisInit()
	_, err := redis.FavoriteList.Del(context.Background(), "1").Result()
	assert.Nil(t, err)
}

func TestRedisGet(t *testing.T) {
	mw.RedisInit()
	val, err := redis.FavoriteList.SMembers(context.Background(), "1").Result()
	fmt.Println(val)
	assert.Nil(t, err)
}
func TestRedisExist(t *testing.T) {
	mw.RedisInit()
	val, err := redis.FavoriteList.Exists(context.Background(), "1").Result()
	fmt.Println(val)
	assert.Nil(t, err)
}
