package service

import "TIKTOK_User/model/vo"

type FriendService interface {
	GetFriendListById(userId, ownerId int64) ([]vo.UserInfo, error)

	// IsFriend 判断user_1 和 user_2 是不是朋友
	IsFriend(userid1, userid2 int64) (bool, error)
}
