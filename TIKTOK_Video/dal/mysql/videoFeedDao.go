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

	return 1, nil
}
func GetCommentCountByID(id int64) (int64, error) {

	return 1, nil
}

func GetVideosByTime(LatestTime int64) ([]model.Video, error) {
	res := make([]model.Video, 2)
	result := DB.Where("publish_time<=?", LatestTime).Order("publish_time desc").Limit(2).Find(&res)

	return res, result.Error
}
func DecrementCommentCount(videoId int64) error {
	video := model.Video{VideoId: videoId}
	result := DB.Model(&video).Update("comment_count", gorm.Expr("comment_count - 1"))
	var err error
	if err = result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return errors.New("减少评论数失败")
	}
	return nil
}

func IncrementCommentCount(videoId int64) error {
	video := model.Video{VideoId: videoId}
	result := DB.Model(&video).Where("video_id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + 1"))
	var err error
	if err = result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return errors.New("添加评论数失败")
	}
	return nil
}
