package model

import (
	"time"
)

type TableVideo struct {
	Id            int64     `column:"video_id"`
	AuthorId      int64     `column:"user_id"`
	PlayUrl       string    `column:"play_url"`
	CoverUrl      string    `column:"cover_url"`
	FavoriteCount int64     `column:"favorite_count"`
	CommentCount  int64     `column:"comment_count"`
	PublishTime   time.Time `column:"publish_time"`
	//IsFavorite    bool   `json:"is_favorite"`
	Title string `json:"title"`
}

func (f *TableVideo) TableName() string {
	return "vms_video"
}
