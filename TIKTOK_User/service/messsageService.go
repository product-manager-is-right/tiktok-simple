package service

type MessageService interface {
	SendMessage(toUserId int64, ownerId int64, content string) error
}
