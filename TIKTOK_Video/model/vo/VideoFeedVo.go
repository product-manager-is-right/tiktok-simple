package vo

import "TIKTOK_Video/model"

type FeedResponse struct {
	StatusCode int32              `json:"status_code"`
	StatusMsg  string             `json:"status_msg,omitempty"`
	VideoList  []model.TableVideo `json:"video_list,omitempty"`
	NextTime   int64              `json:"next_time,omitempty"`
}
