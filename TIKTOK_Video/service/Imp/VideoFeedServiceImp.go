package Imp

import (
	"TIKTOK_Video/dal/mysql"
	"TIKTOK_Video/model"
	"time"
)

type VideoFeedServiceImpl struct {
}

func (VideoService VideoFeedServiceImpl) GetVideoByLatestTime(lastTime time.Time) ([]model.TableVideo, time.Time, error) {
	//videos := make([]model.TableVideo, 0)
	video, err := mysql.GetVideoByTime(lastTime)
	/**
		for _, tmp := range video {
		vi = tmp
		videos = append(videos)
	}
	if err != nil {
		log.Printf("mistake in %v", nil)
	}
	*/
	var timePublic = time.Now()
	//video[len(video)-1].PublishTime
	if len(video) > 0 {
		timePublic = video[len(video)-1].PublishTime
	}
	return video, timePublic, err
}
