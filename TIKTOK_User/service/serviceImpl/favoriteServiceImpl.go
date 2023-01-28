package serviceImpl

import (
	"GoProject/dal/mysql"
	"GoProject/model/vo"
	"errors"
)

type FavoriteServiceImpl struct {
}

func (fsi *FavoriteServiceImpl) GetFavoriteVideosListByUserId(userIdTar, userIdSrc int64) ([]vo.VideoInfo, error) {
	var videoInfos []vo.VideoInfo
	videoIds, err := mysql.GetFavoritesById(userIdTar)

	if err != nil {
		return videoInfos, err
	}
	if len(videoIds) == 0 {
		return videoInfos, errors.New("VideoList is empty")
	}
	videoInfos, err = getVideoInfosByVideoIds(videoIds, userIdSrc, "favorite_query")
	return videoInfos, err

}
func (fsi *FavoriteServiceImpl) GetFavoriteByUserAndVideo(userId, videoId int64) (vo.FavoriteInfo, error) {
	favoriteInfo := vo.FavoriteInfo{}
	favorite, err := mysql.GetFavoriteInfo(1, 25)
	if err != nil {
		return favoriteInfo, errors.New("no record in the database")
	}
	favoriteInfo.Id = favorite.Id
	favoriteInfo.UserId = favorite.UserId
	favoriteInfo.VideoId = favorite.VideoId
	return favoriteInfo, nil
}
