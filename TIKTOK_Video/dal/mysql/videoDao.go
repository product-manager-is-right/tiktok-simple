package mysql

import (
	"TIKTOK_Video/model"
	"errors"
	"gorm.io/gorm"
)

// GetVideoByID 暂时先使用id 后期扩展为时间
func GetVideoByID(id int64) (*model.Video, error) {
	res := &model.Video{}
	if err := DB.Where("video_id = ?", id).
		Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
func GetFavoriteCountByID(id int64) (int64, error) {
	res := &model.Video{}
	if err := DB.Where("video_id = ?", id).
		First(&res).Error; err != nil {
		return 0, err
	}

	return res.FavoriteCount, nil
}
func GetCommentCountByID(id int64) (int64, error) {
	res := &model.Video{}
	if err := DB.Where("video_id = ?", id).
		First(&res).Error; err != nil {
		return 0, err
	}

	return res.CommentCount, nil
}

func GetVideosByTime(LatestTime int64) ([]*model.Video, error) {
	res := make([]*model.Video, 2)
	result := DB.Where("publish_time<?", LatestTime).Order("publish_time desc").Limit(2).Find(&res)

	return res, result.Error
}

func DecrementFavoriteCount(videoId int64) error {
	video := model.Video{VideoId: videoId}
	result := DB.Model(&video).Where("video_id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - 1"))
	var err error
	if err = result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return errors.New("减少点赞数失败")
	}
	return nil
}

func IncrementFavoriteCount(videoId int64) error {
	video := model.Video{VideoId: videoId}
	result := DB.Model(&video).Where("video_id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + 1"))
	var err error
	if err = result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return errors.New("添加点赞数失败")
	}
	return nil
}
