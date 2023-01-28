package serviceImpl

import (
	"GoProject/dal/mysql"
	"GoProject/model/vo"
	"errors"
	"fmt"
)

type FavoriteServiceImpl struct {
}

func (fsi *FavoriteServiceImpl) CreateNewFavorite(userId, videoId int64) (int64, error) {

	favorites, err := mysql.GetFavorite(userId, videoId)
	if err != nil {
		return -1, err
	}
	if len(favorites) > 0 {
		return -1, errors.New("the user has already favored the video")
	}
	favoriteId, err := mysql.CreateNewFavorite(userId, videoId)
	if err != nil {
		return -1, err
	}
	return favoriteId, nil
}
func (fsi *FavoriteServiceImpl) GetFavoriteVideosListByUserId(userIdTar, userIdSrc int64) ([]vo.VideoInfo, error) {
	var videoInfos []vo.VideoInfo
	videoIds, err := mysql.GetFavoritesById(userIdTar)
	if err != nil {
		fmt.Print("userIdTar=", userIdTar)
		//fmt.Print("获取视频id列表失败")
		return videoInfos, err
	}

	if len(videoIds) == 0 {
		return videoInfos, errors.New("获取视频id列表为空")
	}
	videoInfos, err = getVideoInfosByVideoIds(videoIds, userIdSrc, "favorite_query")
	if err != nil {
		fmt.Print("获取视频信息列表失败")
	}
	return videoInfos, err

}

//func (fsi *FavoriteServiceImpl) GetFavoriteByUserAndVideo(userId, videoId int64) (vo.FavoriteInfo, error) {
//	favoriteInfo := vo.FavoriteInfo{}
//	favorite, err := mysql.GetFavoriteInfo(1, 25)
//	if err != nil {
//		return favoriteInfo, errors.New("no record in the database")
//	}
//	favoriteInfo.Id = favorite.Id
//	favoriteInfo.UserId = favorite.UserId
//	favoriteInfo.VideoId = favorite.VideoId
//	return favoriteInfo, nil
//}
