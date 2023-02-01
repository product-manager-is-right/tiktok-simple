package service

import "TIKTOK_Video/model/vo"

type UserService interface {
	GetUserInfoById(queryUserId int64, userId int64) (*vo.UserInfo, error)

	CreateUserByNameAndPassword(username, password string) (int64, error)

	GetUsersInfoByIds(queryUserId []int64, userId int64) (map[int64]*vo.UserInfo, error)
}
