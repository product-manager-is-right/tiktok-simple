package service

import (
	"TIKTOK_Video/service/ServiceImpl"
	"sync"
)

type FavoriteService interface {
	//IsFavorite 判断是否为喜欢接口,假如userId点赞了videoId，返回true,没有返回false
	IsFavorite(userId, videoId int64) (bool, error)
}

var (
	favoriteService FavoriteService

	FavoriteServiceOnce sync.Once
)

// NewFavoriteServiceInstance  单例模式返回service对象
func NewFavoriteServiceInstance() FavoriteService {
	FavoriteServiceOnce.Do(
		func() {
			favoriteService = &ServiceImpl.FavoriteServiceImpl{}
		})
	return favoriteService
}
