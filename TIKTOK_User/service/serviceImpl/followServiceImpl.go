package serviceImpl

import (
	"GoProject/dal/mysql"
	"GoProject/model/vo"
	"gorm.io/gorm"
)

type FollowServiceImpl struct {
}

func (fsi *FollowServiceImpl) GetFollowListById(userId int64) ([]vo.UserInfo, error) {
	//获取关注对象的id数组
	ids, err := mysql.GetFollowingIds(userId)
	if err != nil {
		return []vo.UserInfo{}, err
	}
	// 没得关注者
	if len(ids) == 0 {
		return []vo.UserInfo{}, nil
	}
	// 根据每个id来查询用户信息。
	users := make([]vo.UserInfo, 0, len(ids))
	for i := 0; i < len(ids); i++ {
		isFollow, _ := mysql.GetIsFollow(userId, ids[i])
		if !isFollow {
			continue
		}
		user, err := mysql.GetUserByUserId(ids[i])
		if err == gorm.ErrRecordNotFound {
			continue
		}

		followCnt, _ := mysql.GetFollowCntByUserId(user.Id)

		followerCnt, _ := mysql.GetFollowerCntByUserId(user.Id)
		u := &vo.UserInfo{
			Id: user.Id,
		}
		userInfo.Id = user.Id
		userInfo.Name = user.Username
		userInfo.FollowerCount = followerCnt
		userInfo.FollowCount = followCnt
		userInfo.IsFollow = isFollow
		users = append(users, u)
	}
	return users, nil
}

type FollowerServiceImpl struct {
}

func (fsi *FollowerServiceImpl) GetFollowerListById(userId int64) ([]vo.UserInfo, error) {
	//获取关注对象的id数组
	ids, err := mysql.GetFollowerIds(userId)
	if nil != err {
		return nil, err
	}
	// 没得关注者
	if nil == ids {
		return nil, nil
	}
	// 根据每个id来查询用户信息
	len := len(ids)
	users := make([]vo.UserInfo, len)
	for i := 0; i < len; i++ {
		userInfo := vo.UserInfo{}
		queryUser, err := mysql.GetUserByUserId(ids[i])
		userId = queryUser.Id
		if err != nil {
			return users, err
		}
		followCnt, err := mysql.GetFollowCntByUserId(userId)
		if err != nil {
			return users, err
		}
		followerCnt, err := mysql.GetFollowerCntByUserId(userId)
		if err != nil {
			return users, err
		}
		isFollow, err := mysql.GetIsFollow(userId, ids[i])
		if err != nil {
			return users, err
		}
		userInfo.Id = queryUser.Id
		userInfo.Name = queryUser.Username
		userInfo.FollowerCount = followerCnt
		userInfo.FollowCount = followCnt
		userInfo.IsFollow = isFollow
		users[i] = userInfo
	}
	return users, nil
}
