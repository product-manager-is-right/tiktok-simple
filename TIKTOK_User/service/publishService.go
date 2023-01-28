package service

import "TIKTOK_User/model/vo"

type PublishService interface {
	GetVideoList(userIdTar, userIdSrc int64) ([]vo.VideoInfo, error)
}
