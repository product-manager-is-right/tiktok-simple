package vo

import (
	"TIKTOK_Video/service"
)

type FeedResponse struct {
	StatusCode int32           `json:"status_code"`
	StatusMsg  string          `json:"status_msg,omitempty"`
	VideoList  []service.Video `json:"video_list,omitempty"`
	NextTime   int64           `json:"next_time,omitempty"`
}
