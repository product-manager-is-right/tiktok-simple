package mysql

import (
	"TIKTOK_User/model"
	"time"
)

// message dao层

func CreateMessage(toUserId int64, fromUserId int64, content string) error {
	message := &model.Message{
		UserIdFrom: fromUserId,
		UserIdTo:   toUserId,
		Message:    content,
		CreateTime: time.Now().Unix(),
	}
	result := DB.Create(message)
	return result.Error
}
