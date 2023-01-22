package mysql

import (
	"TIKTOK_Video/model"
	"log"
	"time"
)

// 暂时先使用id 后期扩展为时间
func GetVideoByID(id int64) (*model.TableVideo, error) {
	res := &model.TableVideo{}
	if err := DB.Where("video_id = ?", id).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
func GetFavorCountByID(id int64) (int64, error) {

	return 1, nil
}
func GetCommentCountByID(id int64) (int64, error) {

	return 1, nil
}
func IsFavorByID(id int64) (bool, error) {

	return false, nil
}
func GetVideoByTime(LatestTime time.Time) ([]model.TableVideo, error) {
	res := make([]model.TableVideo, 0)
	log.Print(LatestTime)
	//result := DB.Where("video_id = ?", 1).Limit(1).Find(&res)
	result := DB.Where("publish_time < ?", LatestTime).Order("publish_time desc").Limit(1).Find(&res)
	if result.Error != nil {
		return res, result.Error
	}
	return res, nil
}
