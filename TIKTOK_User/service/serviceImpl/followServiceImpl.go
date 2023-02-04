package serviceImpl

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
	"errors"
	"gorm.io/gorm"
	"log"
	//"fmt"
)

type FollowServiceImpl struct {
}

func (fsi *FollowServiceImpl) CreateNewRelation(userFromId, userToId int64) (int64, error) {
	//var relations []*model.Follow = mysql.Getrelation(usertoid, userfromid)
	relations, err := mysql.GetRelation(userToId, userFromId)
	if err != nil {
		return -1, err
	}
	if len(relations) > 0 {
		return -1, errors.New("the user from id has already followed the usertoid")
	}
	var Cancel int64
	Cancel = 0
	relationId, err := mysql.CreateNewRelation(userToId, userFromId, Cancel)
	if err != nil {
		return -1, err
	}
	return relationId, nil
}
func (fsi *FollowServiceImpl) DeleteRelation(userFromId, userToId int64) error {
	if err := mysql.DeleteRelation(userToId, userFromId); err != nil {
		return err
	}
	return nil
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
	for _, id := range ids {
		userInfo := vo.UserInfo{}
		isFollow, _ := mysql.GetIsFollow(userId, id)
		if !isFollow {
			continue
		}
		user, err := mysql.GetUserByUserId(id)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			log.Print(err)
		}

		followCnt, _ := mysql.GetFollowCntByUserId(user.Id)

		followerCnt, _ := mysql.GetFollowerCntByUserId(user.Id)

		//u := &vo.UserInfo{
		//	Id:            user.Id,
		//	Name:          user.Username,
		//	FollowerCount: followerCnt,
		//	FollowCount:   followCnt,
		//	IsFollow:      isFollow,
		//}
		userInfo.Id = user.Id
		userInfo.Name = user.Username
		userInfo.FollowerCount = followerCnt
		userInfo.FollowCount = followCnt
		userInfo.IsFollow = isFollow
		users = append(users, userInfo)
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
	users := make([]vo.UserInfo, len(ids))
	for index, id := range ids {
		userInfo := vo.UserInfo{}
		queryUser, err := mysql.GetUserByUserId(id)
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
		isFollow, err := mysql.GetIsFollow(userId, id)
		if err != nil {
			return users, err
		}
		userInfo.Id = queryUser.Id
		userInfo.Name = queryUser.Username
		userInfo.FollowerCount = followerCnt
		userInfo.FollowCount = followCnt
		userInfo.IsFollow = isFollow
		users[index] = userInfo
	}
	return users, nil
}
