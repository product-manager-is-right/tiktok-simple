package serviceImpl

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
	"errors"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
}

func (usi *UserServiceImpl) CreateUserByNameAndPassword(username, password string) (int64, error) {
	if len(username) > 32 || len(password) > 32 {
		return -1, errors.New("username or password's length must be < 32")
	}

	// 判断用户名是否存在
	if _, err := mysql.GetUserByUserName(username); err != gorm.ErrRecordNotFound {
		return -1, errors.New("the username has existed")
	}

	userId, err := mysql.CreateUser(username, password)
	if err != nil {
		return -1, err
	}
	return userId, nil
}

func (usi *UserServiceImpl) GetUserInfoById(queryUserId int64, userId int64) (*vo.UserInfo, error) {
	// 调用dal层 ： 根据queryUserId查询username
	userInfo := &vo.UserInfo{}
	queryUser, err := mysql.GetUserByUserId(queryUserId)
	if err != nil {
		// 查询不到queryUser，直接返回空user
		return userInfo, err
	}
	// 调用dal层 ： 根据userId查询关注数和粉丝数，查询失败为0
	followCnt, _ := mysql.GetFollowCntByUserId(queryUserId)

	followerCnt, _ := mysql.GetFollowerCntByUserId(queryUserId)

	//调用dal层 ： 判断主user 是否 关注 queryUser，查询失败为false
	isFollow, _ := mysql.GetIsFollow(queryUserId, userId)

	userInfo.Id = queryUserId
	userInfo.Name = queryUser.Username
	userInfo.FollowerCount = followerCnt
	userInfo.FollowCount = followCnt
	userInfo.IsFollow = isFollow

	return userInfo, nil
}
