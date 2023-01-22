package model

import (
	"TIKTOK_Video/util/vo"
	"time"
)

type TableVideo struct {
	Id            int64   `json:"id"`
	Author        vo.User `json:"user"`
	PlayUrl       string  `json:"play_url"`
	CoverUrl      string  `json:"cover_url"`
	FavoriteCount int64   `json:"favorite_count"`
	CommentCount  int64   `json:"comment_count"`
	PublishTime   time.Time
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

var DemoVideos = []TableVideo{
	{
		Id:            1,
		Author:        vo.DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         "bear",
	},
}

type FeedResponse struct {
	StatusCode int32        `json:"status_code"`
	StatusMsg  string       `json:"status_msg,omitempty"`
	VideoList  []TableVideo `json:"video_list,omitempty"`
	NextTime   int64        `json:"next_time,omitempty"`
}
