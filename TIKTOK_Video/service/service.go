package service

import "TIKTOK_Video/model"

type Video struct {
	model.TableVideo
	//tempId
	IsFavorite bool `json:"is_favorite"`
}
