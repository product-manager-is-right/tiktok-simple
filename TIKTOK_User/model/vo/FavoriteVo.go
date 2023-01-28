package vo

type FavoriteInfo struct {
	Id      int64 `json:"id"`
	UserId  int64 `json:"userId"`
	VideoId int64 `json:"videoId"`
}

type FavoriteListResponse struct {
	Response
	VideoList []VideoInfo `json:"video_list"`
}
type FavoriteInfoResponse struct {
	Response
	FavoriteInfo FavoriteInfo `json:"favorite"`
}
