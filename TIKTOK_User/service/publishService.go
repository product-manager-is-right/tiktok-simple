package service

import "GoProject/model/vo"

type PublishService interface {
	PublishVideo(userId int64, videoData []byte, videoTitle string) error

	GetVideoList(userIdTar, userIdSrc int64) ([]vo.VideoInfo, error)
}
