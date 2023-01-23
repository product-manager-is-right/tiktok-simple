package vo

import (
	"TIKTOK_Video/model"
)

type Video struct {
	model.TableVideo
	//tempId
	IsFavorite bool `json:"is_favorite"`
}

type FeedResponse struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg,omitempty"`
	VideoList  []Video `json:"video_list,omitempty"`
	NextTime   int64   `json:"next_time,omitempty"`
}
