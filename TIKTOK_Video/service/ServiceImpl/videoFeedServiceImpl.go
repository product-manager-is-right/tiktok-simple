package ServiceImpl

import (
	"TIKTOK_Video/dal/mysql"
	"TIKTOK_Video/model"
	"TIKTOK_Video/model/vo"
	"errors"
	"time"
)

type VideoServiceImpl struct {
}

func (vsi *VideoServiceImpl) GetVideoInfosByLatestTime(latestTime int64, userId int64) ([]vo.VideoInfo, int64, error) {
	var videoInfos []vo.VideoInfo
	nextTime := time.Now().Unix()

	videos, err := mysql.GetVideosByTime(latestTime)
	if err != nil {
		return videoInfos, nextTime, err
	}

	if len(videos) == 0 {
		return videoInfos, nextTime, errors.New("video is empty")
	}
	videoInfos = make([]vo.VideoInfo, len(videos))
	videoInfos, err = bindVideoInfo(videoInfos, videos, userId)

	nextTime = videos[len(videos)-1].PublishTime
	return videoInfos, nextTime, err
}

func bindVideoInfo(videoInfos []vo.VideoInfo, videos []model.Video, userId int64) ([]vo.VideoInfo, error) {
	for i, video := range videos {
		videoId := video.Id

		//favoriteCount, err := mysql.GetFavoriteCountByID(videoId)
		//if err != nil {
		//	return videoInfos, err
		//}
		//commentCount, err := mysql.GetCommentCountByID(videoId)
		//if err != nil {
		//	return videoInfos, err
		//}
		// 需要与user通信，应定义到service层
		isFavorite, err := isFavorite(videoId, userId)
		if err != nil {
			return videoInfos, err
		}

		videoInfos[i].Id = videoId
		videoInfos[i].Author = vo.DemoUser
		//videoInfos[i].Author.Id = video.AuthorId
		videoInfos[i].PlayUrl = video.PlayUrl
		videoInfos[i].CoverUrl = video.CoverUrl
		videoInfos[i].FavoriteCount = video.FavoriteCount
		videoInfos[i].CommentCount = video.CommentCount
		videoInfos[i].IsFavorite = isFavorite
		videoInfos[i].Title = video.Title

	}

	return videoInfos, nil
}

// 调用远程接口，判断userId是否喜欢videoId视频
func isFavorite(videoId, userId int64) (bool, error) {
	// TODO : impl
	return false, nil
}
