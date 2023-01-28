package service

import "GoProject/model/vo"

type FavoriteService interface {
	GetFavoriteInfoById(queryUserId int64, userId int64) (vo.UserInfo, error)

	GetFavoriteListByUserId(username, password string) (int64, error)
	GetFavoriteByUserAndVideo(userId int64)
}
