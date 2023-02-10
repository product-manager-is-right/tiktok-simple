package service

import "TIKTOK_User/model/vo"

type MessageService interface {
	SendMessage(toUserId int64, ownerId int64, content string) error
	GetMessage(toUserId int64, ownerId int64) ([]vo.MessageInfo, error)
	GetLatestMessage(toUserId, fromUserId int64) (message string, msgType int64)
}
