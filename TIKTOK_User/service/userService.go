package service

import "GoProject/model/vo"

type UserService interface {
	GetUserInfoById(userId int64) (vo.UserInfo, error)
}
