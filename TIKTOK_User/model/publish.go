package model

import "gorm.io/gorm"

type Publish struct {
	gorm.Model
	Id      int64 `column:"id"`
	UserId  int64 `column:"user_id"`
	VideoId int64 `column:"video_id"`
}

func (p *Publish) TableName() string {
	return "ums_publish_video"
}
