package model

import "time"

type TableVideo struct {
	Id          int64
	AuthorId    int64
	PlayUrl     string
	CoverUrl    string
	PublishTime time.Time
	Title       string //视频名，5.23添加
}
