package Imp

import (
	"TIKTOK_Video/dal/mysql"
	"TIKTOK_Video/model"
	"TIKTOK_Video/service"
	"log"
	"time"
)

type VideoFeedServiceImpl struct {
}

func (VideoService VideoFeedServiceImpl) GetVideoByLatestTime(lastTime time.Time) ([]service.Video, time.Time, error) {

	video, err := mysql.GetVideoByTime(lastTime)
	videos := make([]service.Video, 0)
	var timePublic = time.Now()
	if len(video) <= 0 {
		timePublic = time.Now()

		return videos, timePublic, err
	}
	timePublic = video[len(video)-1].PublishTime
	//videos := make([]service.Video, 0)
	err = VideoService.CopyVideoInfo(&videos, &video)

	/**
		for _, tmp := range video {
		vi = tmp
		videos = append(videos)
	}
	if err != nil {
		log.Printf("mistake in %v", nil)
	}
	*/
	return videos, timePublic, err
}
func (videoService *VideoFeedServiceImpl) CopyVideoInfo(videos *[]service.Video, myVideo *[]model.TableVideo) error {
	for _, video := range *myVideo {
		resId := video.Id
		_, err := mysql.GetVideoByID(resId)
		if err != nil {
			log.Printf("找不到video%v", err)
		}
		favorCount, err := mysql.GetFavorCountByID(resId)
		if err != nil {
			log.Printf("找不到favor%v", err)
		}
		commentCount, err := mysql.GetCommentCountByID(resId)
		if err != nil {
			log.Printf("找不到comment%v", err)
		}
		isfavor, err := mysql.IsFavorByID(resId)
		if err != nil {
			log.Printf("找不到Isfavor%v", err)
		}
		VideoInfo := service.Video{}
		VideoInfo.Id = resId
		VideoInfo.FavoriteCount = favorCount
		VideoInfo.CommentCount = commentCount
		VideoInfo.IsFavorite = isfavor
		VideoInfo.PlayUrl = video.PlayUrl
		VideoInfo.CoverUrl = video.CoverUrl
		VideoInfo.Title = video.Title
		*videos = append(*videos, VideoInfo)
	}
	return nil
}
