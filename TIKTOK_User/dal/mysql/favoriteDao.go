package mysql

import (
	"TIKTOK_User/model"
	"errors"
	"gorm.io/gorm"
)

type FavoriteDao struct {
}

/*
GetIsFavorite
根据userId 和 videoId 判断该用户是否喜欢
*/
func GetIsFavorite(userId, videoId int64) (bool, error) {
	favorite := model.Favorite{}

	if err := DB.Where("user_id = ?", userId).
		Where("video_id = ?", videoId).
		Take(&favorite).Error; err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

// CreateNewFavorite
/*
创建一个Favorite 关联对象进入数据库
*/
func CreateNewFavorite(userId, videoId int64) (int64, error) {
	favorite := model.Favorite{UserId: userId, VideoId: videoId}
	result := DB.Create(&favorite)
	return favorite.Id, result.Error
}

func DeleteFavorite(userId, videoId int64) error {
	Favorite := model.Favorite{UserId: userId, VideoId: videoId}

	result := DB.Where("user_id = ?", userId).Where("video_id = ?", videoId).
		Delete(&Favorite)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("删除失败")
	}
	return nil
}

// GetFavoritesById
/*
通过UserId查询favor关系
*/
func GetFavoritesById(userId int64) ([]int64, error) {
	var res []int64
	if err := DB.Model(model.Favorite{}).Where(map[string]interface{}{"user_id": userId}).Pluck("video_id", &res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

//GetFavorite
/*
查询是否有favorite关联
*/
func GetFavorite(userId, videoId int64) ([]*model.Favorite, error) {
	res := make([]*model.Favorite, 0)
	if err := DB.Where("user_id = ?", userId).Where("video_id = ?", videoId).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
