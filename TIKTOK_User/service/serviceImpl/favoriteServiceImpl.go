package serviceImpl

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
	"errors"
	"log"
)

type FavoriteServiceImpl struct {
}

func (fsi *FavoriteServiceImpl) CreateNewFavorite(userid, videoid int64) (int64, error) {

	favorites, err := mysql.GetFavorite(userid, videoid)
	if err != nil {
		return -1, err
	}
	if len(favorites) > 0 {
		return -1, errors.New("the user has already favored the video")
	}
	favoriteId, err := mysql.CreateNewFavorite(userid, videoid)
	if err != nil {
		return -1, err
	}
	return favoriteId, nil
}
func (fsi *FavoriteServiceImpl) DeleteFavorite(userid, videoid int64) error {
	if err := mysql.DeleteFavorite(userid, videoid); err != nil {
		return err
	}
	return nil
}
func (fsi *FavoriteServiceImpl) GetFavoriteVideosListByUserId(userIdTar, userIdSrc int64) ([]vo.VideoInfo, error) {
	var videoInfos []vo.VideoInfo
	videoIds, err := mysql.GetFavoritesById(userIdTar)
	if err != nil {
		log.Print("userIdTar=", userIdTar)
		//fmt.Print("获取视频id列表失败")
		return nil, err
	}

	if len(videoIds) == 0 {
		return nil, errors.New("获取视频id列表为空")
	}
	videoInfos, err = getVideoInfosByVideoIds(videoIds, userIdSrc, "favorite_query")
	if err != nil {
		log.Print("获取视频信息列表失败")
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
