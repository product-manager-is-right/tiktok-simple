package service

import "GoProject/model/vo"

type PublishService interface {
	GetVideoList(userIdTar, userIdSrc int64) ([]vo.VideoInfo, error)
}
