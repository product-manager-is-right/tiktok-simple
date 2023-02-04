package serviceImpl

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
	"errors"
	"gorm.io/gorm"
	//"fmt"
)

type FollowServiceImpl struct {
}

func (fsi *FollowServiceImpl) CreateNewRelation(userFromId, userToId int64) error {
	relation, err := mysql.GetRelation(userToId, userFromId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	// 数据库没有这条记录，插入
	if err == gorm.ErrRecordNotFound {
		err := mysql.CreateNewRelation(userToId, userFromId)
		return err
	}

	if relation.Cancel == 0 {
		return errors.New("已经关注过了")
	}

	// 数据库已经有这条记录，修改Cancel为0
	if err := mysql.UpdateRelation(userToId, userFromId, 0); err != nil {
		return err
	}

	return nil
}
func (fsi *FollowServiceImpl) DeleteRelation(userFromId, userToId int64) error {
	relation, err := mysql.GetRelation(userToId, userFromId)
	if err == gorm.ErrRecordNotFound {
		return errors.New("没有关注过该用户，无法取关")
	}

	// 目前处于取关状态
	if relation.Cancel == 1 {
		return errors.New("已经取关了")
	}

	if err := mysql.UpdateRelation(userToId, userFromId, 1); err != nil {
		return err
	}

	return nil
}

// GetFollowListById 根据id查询关注列表
func (fsi *FollowServiceImpl) GetFollowListById(userId, ownerId int64) ([]vo.UserInfo, error) {
	//获取关注对象的id数组
	ids, err := mysql.GetFollowingIds(userId)
	if err != nil {
		return []vo.UserInfo{}, err
	}
	// 没关注者
	if len(ids) == 0 {
		return []vo.UserInfo{}, nil
	}
	// 根据每个id来查询用户信息
	users := make([]vo.UserInfo, 0, len(ids))
	for _, id := range ids {
		user, err := mysql.GetUserByUserId(id)
		if err != nil {
			continue
		}

		followCnt, _ := mysql.GetFollowCntByUserId(id)

		followerCnt, _ := mysql.GetFollowerCntByUserId(id)

		isFollow, _ := mysql.GetIsFollow(id, ownerId)

		u := vo.UserInfo{
			Id:            id,
			Name:          user.Username,
			FollowerCount: followerCnt,
			FollowCount:   followCnt,
			IsFollow:      isFollow,
		}
		users = append(users, u)
	}
	return users, nil
}

type FollowerServiceImpl struct {
}

// GetFollowerListById 根据id查询粉丝列表
func (fsi *FollowerServiceImpl) GetFollowerListById(userId, ownerId int64) ([]vo.UserInfo, error) {
	//获取关注对象的id数组
	ids, err := mysql.GetFollowerIds(userId)
	if err != nil {
		return []vo.UserInfo{}, err
	}
	// 没粉丝
	if len(ids) == 0 {
		return []vo.UserInfo{}, nil
	}
	// 根据每个id来查询用户信息
	users := make([]vo.UserInfo, 0, len(ids))
	for _, id := range ids {
		user, err := mysql.GetUserByUserId(id)
		if err != nil {
			continue
		}

		followCnt, _ := mysql.GetFollowCntByUserId(id)

		followerCnt, _ := mysql.GetFollowerCntByUserId(id)

		isFollow, _ := mysql.GetIsFollow(id, ownerId)

		u := vo.UserInfo{
			Id:            id,
			Name:          user.Username,
			FollowerCount: followerCnt,
			FollowCount:   followCnt,
			IsFollow:      isFollow,
		}
		users = append(users, u)
	}
	return users, nil
}
