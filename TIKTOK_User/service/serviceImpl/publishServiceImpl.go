package serviceImpl

import (
	"GoProject/dal/mysql"
	"GoProject/model/vo"
	"errors"
	"fmt"
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

@ params

	userIdTar: 请求查询的userId
	userIdSrc: 登录userId
*/
func (psi *PublishServiceImpl) GetVideoList(userIdTar, userIdSrc int64) ([]vo.VideoInfo, error) {
	// 调用Dao层，查找userIdTar用户的所有videoIds
	var videoInfos []vo.VideoInfo
	videoIds, err := mysql.GetPublishVideoIdsById(userIdTar)
	if err != nil {
		return videoInfos, err
	}
	if len(videoIds) == 0 {
		return videoInfos, errors.New("VideoList is empty")
	}
	videoInfos, err = getVideoInfosByVideoIds(videoIds, userIdSrc, "publish_query")
	return videoInfos, err
}

/*
通过videoIds封装VideoInfo，发布视频列表和喜欢列表可用
@ params

	videoIds: 视频ids
	userId: 登录的userId，用于查询登录账号与<视频(是否喜欢)和用户(是否关注)>的关系
*/
func getVideoInfosByVideoIds(videoIds []int64, userId int64, mode string) ([]vo.VideoInfo, error) {
	videoInfos, err := remoteGetVideoInfoCall(videoIds)
	if err != nil {
		fmt.Print("远程调用视频信息失败")
		return videoInfos, err
	}

	// 处理每个VideoInfo中的User和is_favorite
	usi := UserServiceImpl{}

	// 发布视频模式，发布的用户信息一致，无需重复查询
	if mode == "publish_query" {
		user, _ := usi.GetUserInfoById(videoInfos[0].Author.Id, userId)
		for i, videoInfo := range videoInfos {
			videoInfos[i].Author = user

			if isFavorite, _ := mysql.GetIsFavorite(userId, videoInfo.Id); isFavorite {
				videoInfos[i].IsFavorite = true
			}
		}
	} else {
		for i, videoInfo := range videoInfos {
			user, _ := usi.GetUserInfoById(videoInfo.Author.Id, userId)
			videoInfos[i].Author = user

			//if isFavorite, _ := mysql.GetIsFavorite(userId, videoInfo.Id); isFavorite {
			videoInfos[i].IsFavorite = true
			//}
		}
	}

	return videoInfos, nil
}

/*
远程调用Video模块，获取每个Video的具体信息
*/
func remoteGetVideoInfoCall(videoIds []int64) ([]vo.VideoInfo, error) {
	// TODO : impl
	res := make([]vo.VideoInfo, 0)
	for i := range videoIds {
		res = append(res, vo.VideoInfo{Id: videoIds[i]})
	}
	return res, nil
}
