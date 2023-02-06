package serviceImpl

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
)

type FriendServiceImpl struct {
}

// GetFriendListById 根据id获取好友列表
func (fsi *FriendServiceImpl) GetFriendListById(userId, ownerId int64) ([]vo.UserInfo, error) {
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
	for _, id := range ids {
		f, _ := mysql.GetIsFollow(userId, id)
		if !f {
			continue
		}
		user, err := mysql.GetUserByUserId(id)
		if err != nil {
			continue
		}

		followCnt, _ := mysql.GetFollowCntByUserId(user.Id)

		followerCnt, _ := mysql.GetFollowerCntByUserId(user.Id)

		isFollow, _ := mysql.GetIsFollow(id, ownerId)

		u := vo.UserInfo{
			Id:            user.Id,
			Name:          user.Username,
			FollowerCount: followerCnt,
			FollowCount:   followCnt,
			IsFollow:      isFollow,
		}

		users = append(users, u)
	}

	return users, nil
}

func (fsi *FriendServiceImpl) IsFriend(userid1, userid2 int64) (bool, error) {
	a, err := mysql.GetIsFollow(userid1, userid2)
	if err != nil {
		return false, err
	}
	if a == false {
		return false, nil
	}
	b, err := mysql.GetIsFollow(userid2, userid1)
	if err != nil {
		return false, err
	}

	return b, err
}
