package service

import "TIKTOK_User/model/vo"

type UserService interface {
	GetUserInfoById(queryUserId int64, userId int64) (*vo.UserInfo, error)

	CreateUserByNameAndPassword(username, password string) (int64, error)

	GetUsersInfoByIds(queryUserId []int64, userId int64) ([]*vo.UserInfo, error)
}
