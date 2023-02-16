package serviceImpl

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
	"TIKTOK_User/mw/rabbitMQ/producer"
	"TIKTOK_User/mw/redis"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"strconv"
	//"fmt"
)

type FollowServiceImpl struct {
}

func (fsi *FollowServiceImpl) CreateRelation(userFromId, userToId int64) error {
	_, err := mysql.GetRelation(userToId, userFromId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	// 数据库没有这条记录，插入
	if err == gorm.ErrRecordNotFound {
		//使用rabbitMQ
		err = producer.SendFollowMessage(userToId, userFromId, 1)
		if err != nil {
			log.Print("启动rabbitMQ失败，使用Mysql直接处理数据")
			if err := mysql.CreateRelation(userToId, userFromId); err != nil {
				return err
			}
			// follow数据库已经改变，关注列表和粉丝列表都要删除对应的key， 重试机制保证删除
			strUserFromId := strconv.FormatInt(userFromId, 10)
			strUserToId := strconv.FormatInt(userToId, 10)
			for i := 0; i < redis.RetryTime; i++ {
				if _, err := redis.FollowList.Del(context.Background(), strUserFromId).Result(); err == nil {
					break
				}
			}
			for i := 0; i < redis.RetryTime; i++ {
				if _, err := redis.FollowerList.Del(context.Background(), strUserToId).Result(); err == nil {
					break
				}
			}
		}
		return nil
	}

	err = errors.New("已经关注了")
	return err
}

func (fsi *FollowServiceImpl) DeleteRelation(userFromId, userToId int64) error {
	_, err := mysql.GetRelation(userToId, userFromId)
	if err == gorm.ErrRecordNotFound {
		return errors.New("没有关注过该用户，无法取关")
	}

	// 数据库有这条记录，删除
	// 使用rabbitMQ
	err = producer.SendFollowMessage(userToId, userFromId, 0)
	if err != nil {
		log.Print("启动rabbitMQ失败，使用Mysql直接处理数据")
		if err := mysql.DeleteRelation(userToId, userFromId); err != nil {
			return err
		}
		// follow数据库已经改变，关注列表和粉丝列表都要删除对应的key， 重试机制保证删除
		strUserFromId := strconv.FormatInt(userFromId, 10)
		strUserToId := strconv.FormatInt(userToId, 10)
		for i := 0; i < redis.RetryTime; i++ {
			if _, err := redis.FollowList.Del(context.Background(), strUserFromId).Result(); err == nil {
				break
			}
		}
		for i := 0; i < redis.RetryTime; i++ {
			if _, err := redis.FollowerList.Del(context.Background(), strUserToId).Result(); err == nil {
				break
			}
		}
	}
	return nil
}

// GetFollowListById 根据id查询关注列表
func (fsi *FollowServiceImpl) GetFollowListById(userId, ownerId int64) ([]vo.UserInfo, error) {
	ids := make([]int64, 0, 10)

	// 先从redis中查找
	strUserId := strconv.FormatInt(userId, 10)
	if n, err := redis.FollowList.Exists(context.Background(), strUserId).Result(); err == nil && n > 0 {
		// 缓存命中
		vs, err := redis.FollowList.SMembers(context.Background(), strUserId).Result()
		if err != nil {
			return nil, err
		}
		// 转换str->int64
		for _, v := range vs {
			r, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, errors.New("redis中存储非法")
			}
			ids = append(ids, r)
		}
	} else {
		// 缓存未命中，查询数据库
		ids, err = mysql.GetFollowingIds(userId)
		if err != nil {
			return nil, errors.New("数据库查询失败")
		}
		// 转换int64->str
		followIds := make([]string, len(ids))
		for i, v := range ids {
			followIds[i] = strconv.FormatInt(v, 10)
		}
		// 存入redis，不需要处理异常
		redis.FollowList.SAdd(context.Background(), strUserId, followIds)
		// 设置过期时间，兜底方案
		if _, err := redis.FollowList.Expire(context.Background(), strUserId, redis.SetExpiredTime()).Result(); err != nil {
			// 设置失败，删除该key
			redis.FollowList.Del(context.Background(), strUserId)
		}
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
	ids := make([]int64, 0, 10)

	// 先从redis中查找
	strUserId := strconv.FormatInt(userId, 10)
	if n, err := redis.FollowerList.Exists(context.Background(), strUserId).Result(); err == nil && n > 0 {
		// 缓存命中
		vs, err := redis.FollowerList.SMembers(context.Background(), strUserId).Result()
		if err != nil {
			return nil, err
		}
		// 转换str->int64
		for _, v := range vs {
			r, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, errors.New("redis中存储非法")
			}
			ids = append(ids, r)
		}
	} else {
		// 缓存未命中，查询数据库
		ids, err = mysql.GetFollowerIds(userId)
		if err != nil {
			return nil, errors.New("数据库查询失败")
		}
		// 转换int64->str
		followerIds := make([]string, len(ids))
		for i, v := range ids {
			followerIds[i] = strconv.FormatInt(v, 10)
		}
		// 存入redis，不需要处理异常
		redis.FollowerList.SAdd(context.Background(), strUserId, followerIds)
		// 设置过期时间，兜底方案，设置随机的过期时间，防止缓存雪崩
		if _, err := redis.FollowerList.Expire(context.Background(), strUserId, redis.SetExpiredTime()).Result(); err != nil {
			// 设置失败，删除该key
			redis.FollowerList.Del(context.Background(), strUserId)
		}
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
