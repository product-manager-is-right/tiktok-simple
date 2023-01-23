package service

import (
	"TIKTOK_Video/model/vo"
)

type VideoService interface {
	GetVideoByLatestTime(lastTime int64) ([]vo.VideoInfo, int64, error)
}
