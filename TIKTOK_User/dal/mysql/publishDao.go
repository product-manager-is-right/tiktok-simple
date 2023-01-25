package mysql

/*
CreatePublishVideo
根据UserId 和 VideoId，在ums_publish_video添加一行数据
*/
func CreatePublishVideo(UserId, VideoId int64) error {
	// TODO : impl
	return nil
}

/*
GetPublishVideoIdsById
根据UserId查找所有发布的VideoIds
*/
func GetPublishVideoIdsById(UserId int64) ([]int64, error) {
	// TODO : impl
	return []int64{}, nil
}
