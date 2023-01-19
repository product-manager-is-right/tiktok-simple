package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int64  `column:"id"`
	Name     string `column:"username"`
	Password string `column:"password"`
}

func (u *User) TableName() string {
	return "ums_user"
}
