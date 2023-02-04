package service

import "TIKTOK_User/model/vo"

type FriendService interface {
	GetFriendListById(userId, ownerId int64) ([]vo.UserInfo, error)
}
