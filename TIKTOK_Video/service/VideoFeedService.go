package service

import (
	"TIKTOK_Video/model"
	"time"
)

type VideoService interface {
	//GetVideoInfoById(video int64) (model.TableVideo, error)

	GetVideoByLatestTime(lastTime time.Time) ([]model.TableVideo, time.Time, error)
}
