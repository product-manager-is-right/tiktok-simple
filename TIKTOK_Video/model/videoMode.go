package model

import (
	"time"
)

type TableVideo struct {
	Id            int64     `json:"video_id"`
	AuthorId      int64     `json:"user_id"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	PublishTime   time.Time `json:"publish_time"`
	//IsFavorite    bool   `json:"is_favorite"`
	Title string `json:"title"`
}

func (f *TableVideo) TableName() string {
	return "vms_video"
}
