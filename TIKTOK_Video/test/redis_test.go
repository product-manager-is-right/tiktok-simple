package test

import (
	"TIKTOK_Video/mw"
	"TIKTOK_Video/mw/redis"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisAdd(t *testing.T) {
	mw.RedisInit()
	_, err := redis.CommentList.SAdd(context.Background(), "1", "1", "2", "3").Result()
	assert.Nil(t, err)
}

func TestRedisDel(t *testing.T) {
	mw.RedisInit()
	_, err := redis.CommentList.Del(context.Background(), "1").Result()
	assert.Nil(t, err)
}

func TestRedisGet(t *testing.T) {
	mw.RedisInit()
	val, err := redis.CommentList.SMembers(context.Background(), "1").Result()
	fmt.Println(val)
	assert.Nil(t, err)
}
func TestRedisExist(t *testing.T) {
	mw.RedisInit()
	val, err := redis.CommentList.Exists(context.Background(), "1").Result()
	fmt.Println(val)
	assert.Nil(t, err)
}
