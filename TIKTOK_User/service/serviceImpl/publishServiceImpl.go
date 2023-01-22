package serviceImpl

import (
	"GoProject/dal/mysql"
	"GoProject/model/vo"
	"errors"
)

type PublishServiceImpl struct {
}

func (psi *PublishServiceImpl) PublishVideo(userId int64, videoData []byte, videoTitle string) error {
	// 远程调用video接口
	videoId, err := remoteCreateVideoCall(userId, videoData, videoTitle)
	if err != nil {
		return err
	}

	// 调用Dao层，存入ums_publish_video
	if err = mysql.CreatePublishVideo(userId, videoId); err != nil {
		return err
	}
	return nil
}

/*
remoteVideoCall
发起远程调用视频模块，存储video，返回videoId
@ return videoId
*/
func remoteCreateVideoCall(userId int64, videoData []byte, videoTitle string) (int64, error) {
	// TODO : impl
	return 0, nil
}

/*
GetVideoList
获取userId发布的视频列表
*/
func (psi *PublishServiceImpl) GetVideoList(userIdTar, userIdSrc int64) ([]vo.VideoInfo, error) {
	// 调用Dao层，查找userIdTar用户的所有videoIds
	var videoInfos []vo.VideoInfo
	videoIds, err := mysql.GetVideoIdsById(userIdTar)
	if err != nil {
		return videoInfos, err
	}
	if len(videoIds) == 0 {
		return videoInfos, errors.New("VideoList is empty")
	}
	videoInfos, err = remoteGetVideoInfo(videoIds)
	if err != nil {
		return videoInfos, err
	}

	// 处理每个VideoInfo中的User和is_favorite
	usi := UserServiceImpl{}
	user, err := usi.GetUserInfoById(userIdTar, userIdSrc)
	for i, videoInfo := range videoInfos {
		videoInfos[i].Author = user
		if isFavorite, _ := mysql.GetIsFavorite(userIdSrc, videoInfo.Id); isFavorite {
			videoInfos[i].IsFavorite = true
		}
	}
	return videoInfos, err
}

/*
远程调用Video模块，获取每个Video的具体信息
*/
func remoteGetVideoInfo(videoIds []int64) ([]vo.VideoInfo, error) {
	// TODO : impl
	return []vo.VideoInfo{}, nil
}
