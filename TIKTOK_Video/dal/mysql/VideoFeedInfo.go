package mysql

import (
	"TIKTOK_Video/model"
	"time"
)

// 暂时先使用id 后期扩展为时间
func GetVideoByID(id int64) ([]*model.TableVideo, error) {
	res := make([]*model.TableVideo, 0)
	if err := DB.Where("video_id = ?", id).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
func GetVideoByTime(time time.Time) ([]*model.TableVideo, error) {
	res := make([]*model.TableVideo, 0)
	if err := DB.Where("publish_time<?", time).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
