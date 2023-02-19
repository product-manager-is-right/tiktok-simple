package serviceImpl

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
)

type FriendServiceImpl struct {
}

// GetFriendListById 根据id获取好友列表
func (fsi *FriendServiceImpl) GetFriendListById(userId, ownerId int64) ([]vo.FriendUserInfo, error) {
	//获取关注对象的id数组
	ids, err := mysql.GetFollowingIds(userId)
	if err != nil {
		return []vo.FriendUserInfo{}, err
	}
	// 没得关注者
	if len(ids) == 0 {
		return []vo.FriendUserInfo{}, nil
	}
	// 根据每个id来查询用户信息。
	friends := make([]vo.FriendUserInfo, 0, len(ids))
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

		msi := MessageServiceImpl{}
		message, msgType := msi.GetLatestMessage(id, userId)

		fu := vo.FriendUserInfo{
			UserInfo: vo.UserInfo{
				Id:            id,
				Name:          user.Username,
				FollowerCount: followerCnt,
				FollowCount:   followCnt,
				IsFollow:      isFollow,
				Avatar:        "http://120.25.2.146:9000/tikpic/head1.jpg",
			},
			Message: message,
			MsgType: msgType,
		}

		friends = append(friends, fu)
	}

	return friends, nil
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
