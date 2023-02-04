package service

import (
	"TIKTOK_User/model/vo"
	"TIKTOK_User/service/serviceImpl"
	"sync"
)

type FollowService interface {
	CreateNewRelation(userFromId, userToId int64) (int64, error)
	DeleteRelation(userFromId, userToId int64) error
	GetFollowListById(userId int64) ([]vo.UserInfo, error)
}

var (
	followService FollowService

	followServiceOnce sync.Once
)

// NewFollowServiceInstance  单例模式返回service对象
func NewFollowServiceInstance() FollowService {
	followServiceOnce.Do(
		func() {
			followService = &serviceImpl.FollowServiceImpl{}
		})
	return followService
}
