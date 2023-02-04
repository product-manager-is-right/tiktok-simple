package service

import "TIKTOK_User/model/vo"

type FavoriteService interface {
	CreateNewFavorite(userId, videoId int64) error
	DeleteFavorite(userId, videoId int64) error
	GetFavoriteVideosListByUserId(userIdTar, userIdSrc int64) ([]vo.VideoInfo, error)
	IsFavorite(userId, videoId int64) (bool, error)
}
