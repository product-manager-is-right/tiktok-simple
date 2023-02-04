package service

import (
	"TIKTOK_User/model/vo"
)

type FollowService interface {
	CreateNewRelation(userFromId, userToId int64) error
	DeleteRelation(userFromId, userToId int64) error
	GetFollowListById(userId, ownerId int64) ([]vo.UserInfo, error)
}

type FollowerService interface {
	GetFollowerListById(userId, ownerId int64) ([]vo.UserInfo, error)
}
