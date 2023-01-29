package mysql

import (
	"GoProject/model"
	"fmt"
)

/*
GetIsFavorite
根据userId 和 videoId 判断该用户是否喜欢
*/
func GetIsFavorite(userId, videoId int64) (bool, error) {
	// TODO : impl
	return true, nil
}
func CreateNewFavorite(userId, videoId int64) (int64, error) {
	favorite := model.Favorite{UserId: userId, VideoId: videoId}
	result := DB.Create(&favorite)
	return favorite.Id, result.Error
}
func GetFavoritesById(userId int64) ([]int64, error) {
	var res []int64
	err := DB.Model(model.Favorite{}).Where(map[string]interface{}{"user_id": userId}).Pluck("video_id", &res).Error
	if err != nil {
		fmt.Print(err)
	}

	return res, nil
}
func GetFavorite(userId, videoId int64) ([]*model.Favorite, error) {
	res := make([]*model.Favorite, 0)
	if err := DB.Where("user_id = ?", userId).Where("video_id = ?", videoId).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
