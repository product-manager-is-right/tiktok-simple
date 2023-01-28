package service

type FavoriteService interface {
	CreateNewFavorite(userId, videoId int64) (int64, error)
	GetFavoriteListByUserId(username, password string) (int64, error)
}
