package service

import (
	"TIKTOK_User/service/serviceImpl"
	"sync"
)

type FavoriteService interface {
	CreateNewFavorite(userid, videoid int64) (int64, error)
	DeleteFavorite(userid, videoid int64) error
	//GetFavoriteListByUserId(username, password string) (int64, error)
}

var (
	favoriteservice FavoriteService

	FavoriteServiceOnce sync.Once
)

// NewCommentServiceInstance  单例模式返回service对象
func NewFavoriteServiceInstance() FavoriteService {
	FavoriteServiceOnce.Do(
		func() {
			favoriteservice = &serviceImpl.FavoriteServiceImpl{}
		})
	return favoriteservice
}
