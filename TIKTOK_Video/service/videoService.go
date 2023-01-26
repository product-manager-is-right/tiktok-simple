package service

import (
	"TIKTOK_Video/model/vo"
	"mime/multipart"
)

type VideoService interface {
	PublishVideo(userId int64, videoData *multipart.FileHeader, videoTitle string) error

	GetVideoInfosByLatestTime(lastTime int64, userName string) ([]vo.VideoInfo, int64, error)
}
