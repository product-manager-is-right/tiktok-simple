package model

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	Id         int64 `column:"id"`
	UserIdFrom int64 `column:"user_id_from"`
	UserIdTo   int64 `column:"user_id_to"`
	Cancel     int   `column:"cancel"`
}

func (f *Friend) TableName() string {
	return "ums_friend"
}
