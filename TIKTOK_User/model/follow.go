package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	Id         int64 `column:"id"`
	UserIdFrom int64 `column:"user_id_from"`
	UserIdTo   int64 `column:"user_id_to"`
	Cancel     int64 `column:"cancel"`
}

func (f *Follow) TableName() string {
	return "ums_follow"
}
