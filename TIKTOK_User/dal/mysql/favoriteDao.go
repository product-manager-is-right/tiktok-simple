package mysql

import (
	"GoProject/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type FavoriteDao struct {
}

/*
GetIsFavorite
根据userId 和 videoId 判断该用户是否喜欢
*/
func GetIsFavorite(userid, videoid int64) (bool, error) {
	favorite := model.Favorite{}

	if err := DB.Where("user_id = ?", userid).
		Where("video_id = ?", videoid).
		Take(&favorite).Error; err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
func CreateNewFavorite(userid, videoid int64) (int64, error) {
	favorite := model.Favorite{UserId: userid, VideoId: videoid}
	result := DB.Create(&favorite)
	return favorite.Id, result.Error
}

func DeleteFavorite(userid, videoid int64) error {
	Favorite := model.Favorite{UserId: userid, VideoId: videoid}

	result := DB.Where("user_id = ?", userid).Where("video_id = ?", videoid).
		Delete(&Favorite)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("删除失败")
	}
	return nil
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
