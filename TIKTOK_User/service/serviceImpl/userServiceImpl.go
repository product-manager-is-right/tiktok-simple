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

	workCount, _ := mysql.GetPublishVideoIdsById(queryUserId)

	favorCount, _ := mysql.GetFavoritesById(queryUserId)

	userInfo.Id = queryUserId
	userInfo.Name = queryUser.Username
	userInfo.FollowerCount = followerCnt
	userInfo.FollowCount = followCnt
	userInfo.IsFollow = isFollow
	userInfo.WorkCount = int64(len(workCount))
	userInfo.FavoriteCount = int64(len(favorCount))
	userInfo.Signature = "speak less, do more"
	userInfo.Background = "http://120.25.2.146:9000/tikpic/backgg.jpg"
	userInfo.Avatar = "http://120.25.2.146:9000/tikpic/head.jpg"
	return userInfo, nil
}

func (usi *UserServiceImpl) GetUsersInfoByIds(queryUserId []int64, userId int64) ([]*vo.UserInfo, error) {
	users, err := mysql.GetUserByIds(queryUserId)
	if err != nil {
		return nil, err
	}
	userInfo := make([]*vo.UserInfo, len(users))
	for idx, user := range users {
		if user == nil {
			continue
		}
		info := &vo.UserInfo{}
		// 调用dal层 ： 根据userId查询关注数和粉丝数，查询失败为0
		followCnt, _ := mysql.GetFollowCntByUserId(user.Id)

		followerCnt, _ := mysql.GetFollowerCntByUserId(user.Id)

		//调用dal层 ： 判断主user 是否 关注 queryUser，查询失败为false
		isFollow, _ := mysql.GetIsFollow(user.Id, userId)
		workCount, _ := mysql.GetPublishVideoIdsById(user.Id)

		favorCount, _ := mysql.GetFavoritesById(user.Id)

		info.Id = user.Id
		info.Name = user.Username
		info.FollowerCount = followerCnt
		info.FollowCount = followCnt
		info.IsFollow = isFollow
		info.WorkCount = int64(len(workCount))
		info.FavoriteCount = int64(len(favorCount))
		info.Signature = "speak less, do more"
		info.Background = "http://120.25.2.146:9000/tikpic/backg.jpg"
		info.Avatar = "http://120.25.2.146:9000/tikpic/head1.jpg"

		userInfo[idx] = info
	}
	return userInfo, nil
}
