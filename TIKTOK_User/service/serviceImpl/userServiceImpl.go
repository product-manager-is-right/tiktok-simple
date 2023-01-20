package serviceImpl

import (
	"GoProject/dal/mysql"
	"GoProject/model/vo"
	"errors"
)

type UserServiceImpl struct {
}

func (usi *UserServiceImpl) CreateUserByNameAndPassword(username, password string) (int64, error) {
	if len(username) > 32 || len(password) > 32 {
		return -1, errors.New("username or password's length > 32")
	}
	users, err := mysql.GetUserByUserName(username)
	if err != nil {
		return -1, err
	}
	if len(users) > 0 {
		return -1, errors.New("this username has exist")
	}
	userId, err := mysql.CreateUser(username, password)
	if err != nil {
		return -1, err
	}
	return userId, nil
}

func (usi *UserServiceImpl) GetUserInfoById(queryUserId int64, userId int64) (vo.UserInfo, error) {
	// 调用dal层 ： 根据userId查询username
	userInfo := vo.UserInfo{}
	queryUser, err := mysql.GetUserByUserId(queryUserId)
	if err != nil {
		return userInfo, err
	}

	// 调用dal层 ： 根据userId查询关注数和粉丝数
	followCnt, err := mysql.GetFollowCntByUserId(queryUserId)
	if err != nil {
		return userInfo, err
	}
	followerCnt, err := mysql.GetFollowerCntByUserId(queryUserId)
	if err != nil {
		return userInfo, err
	}

	//调用dal层 ： 判断userId 是否 关注 queryUserId
	isFollow, err := mysql.GetIsFollow(queryUserId, userId)
	if err != nil {
		return userInfo, err
	}

	userInfo.Id = queryUserId
	userInfo.Name = queryUser.Username
	userInfo.FollowerCount = followerCnt
	userInfo.FollowCount = followCnt
	userInfo.IsFollow = isFollow

	return userInfo, nil
}
