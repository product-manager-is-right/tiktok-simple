package model

type Favorite struct {
	Id      int64  `column:"id"`
	UserId  string `column:"user_id"`
	VideoId string `column:"video_id"`
}

func (f *Favorite) TableName() string {
	return "ums_favorite_video"
}
