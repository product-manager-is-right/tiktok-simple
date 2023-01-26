package mysql

import (
	"GoProject/model"
	"time"
)

/*
CreatePublishVideo
根据UserId 和 VideoId，在ums_publish_video添加一行数据
*/
func CreatePublishVideo(UserId, VideoId int64) error {
	// TODO : impl
	video := model.Publish{UserId: UserId, VideoId: VideoId}
	result := DB.Create(&video)
	return result.Error
}

/*
CreateVideo
根据UserId 和 VideoId url，在vms中添加一行数据
*/
func CreateVideo(UserId int64, playUrl string, coverUrl string, title string) (int64, error) {
	// TODO : impl
	video := model.Video{UserId: UserId, PlayUrl: playUrl, CoverUrl: coverUrl, Title: title, PublishTime: time.Now().Unix()}
	result := DBV.Create(&video)
	return video.Id, result.Error
}

/*
GetPublishVideoIdsById
根据UserId查找所有发布的VideoIds
*/
func GetPublishVideoIdsById(UserId int64) ([]int64, error) {
	// TODO : impl
	return []int64{}, nil
}
