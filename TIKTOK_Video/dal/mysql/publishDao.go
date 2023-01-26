package mysql

import (
	"TIKTOK_Video/model"
	"time"
)

/*
CreateVideo
根据UserId 和 VideoId url，在vms中添加一行数据
*/
func CreateVideo(UserId int64, playUrl string, coverUrl string, title string) (int64, error) {
	// TODO : impl
	video := model.Video{UserId: UserId, PlayUrl: playUrl, CoverUrl: coverUrl, Title: title, PublishTime: time.Now().Unix()}
	result := DB.Create(&video)
	return video.VideoId, result.Error
}

/*
GetPublishVideoIdsById
根据UserId查找所有发布的VideoIds
*/
func GetPublishVideoIdsById(UserId int64) ([]int64, error) {
	// TODO : impl
	return []int64{}, nil
}
