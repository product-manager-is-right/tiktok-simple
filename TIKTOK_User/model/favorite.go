package model

type Favorite struct {
	Id      int64 `column:"id"`
	UserId  int64 `column:"user_id"`
	VideoId int64 `column:"video_id"`
	Cancel  int64 `column:"cancel"`
}

func (f *Favorite) TableName() string {
	return "ums_favorite_video"
}
