package mysql

import (
	"GoProject/model"
)

/*
GetIsFavorite
根据userId 和 videoId 判断该用户是否喜欢
*/
func GetIsFavorite(userId, videoId int64) (bool, error) {
	// TODO : impl
	return true, nil
}

func GetFavoritesById(userId int64) ([]int64, error) {
	res := make([]int64, 0)
	if err := DB.Where("user_id = ?", userId).
		Pluck("video_id", &res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
func GetFavoriteInfo(userId, videoId int64) (*model.Favorite, error) {
	res := &model.Favorite{}
	if err := DB.Where("user_id = ?", userId).
		First(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
