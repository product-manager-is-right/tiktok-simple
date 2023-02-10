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

func GetMessage(toUserId int64, ownerId int64) ([]model.Message, error) {
	var res1 []model.Message
	if err := DB.Model(model.Message{}).Where("user_id_to = ?", toUserId).Where("user_id_from = ?", ownerId).Find(&res1).Error; err != nil {
		return nil, err
	}
	var res2 []model.Message
	if err := DB.Model(model.Message{}).Where("user_id_to = ?", ownerId).Where("user_id_from = ?", toUserId).Find(&res2).Error; err != nil {
		return nil, err
	}
	res1 = append(res1, res2...)

	return res1, nil
}

func GetLatestMessage(toUserId, fromUserId int64) model.Message {
	var res1 model.Message
	DB.Model(model.Message{}).Where("user_id_to = ?", toUserId).
		Where("user_id_from = ?", fromUserId).Order("create_time desc").Limit(1).Take(&res1)

	var res2 model.Message
	DB.Model(model.Message{}).Where("user_id_to = ?", fromUserId).
		Where("user_id_from = ?", toUserId).Order("create_time desc").Limit(1).Take(&res2)

	if res1.CreateTime > res2.CreateTime {
		return res1
	}
	return res2
}
