package serviceImpl

import (
	"GoProject/dal/mysql"
	"GoProject/model/vo"
)

type UserServiceImpl struct {
}

func (usi *UserServiceImpl) GetUserInfoById(userId int64) (vo.UserInfo, error) {
	// 调用dal层 ： 根据userId查询username
	userInfo := vo.UserInfo{}
	user, err := mysql.GetUserByUserId(userId)
	if err != nil {
		return userInfo, err
	}
	// 调用dal层 ： 根据userId查询关注数和粉丝数
	followCnt, err := mysql.GetFollowCntByUserId(userId)
	if err != nil {
		return userInfo, err
	}
	followerCnt, err := mysql.GetFollowerCntByUserId(userId)
	if err != nil {
		return userInfo, err
	}
	userInfo.Id = userId
	userInfo.Name = user.Name
	userInfo.FollowerCount = followerCnt
	userInfo.FollowCount = followCnt
	userInfo.IsFollow = false

	return userInfo, nil
}
