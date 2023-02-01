package mysql

import (
	"GoProject/model"
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
GetPublishVideoIdsById
根据UserId查找所有发布的VideoIds
*/
func GetPublishVideoIdsById(UserId int64) ([]int64, error) {
	res := make([]*model.Publish, 0)
	result := DB.Where("user_id = ?", UserId).Find(&res)
	videoIds := make([]int64, 0)
	for _, publish := range res {
		videoIds = append(videoIds, publish.VideoId)
	}
	// TODO : impl
	return videoIds, result.Error
}
