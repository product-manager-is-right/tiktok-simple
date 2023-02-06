package serviceImpl

import (
	"TIKTOK_User/dal/mysql"
	"errors"
)

type MessageServiceImpl struct {
}

var ErrNotFriend = errors.New("not friend relation")

func (msi *MessageServiceImpl) SendMessage(toUserId int64, ownerId int64, content string) error {
	// TODO : 判断toUserId是否为已注册的合法id (应由关注模块约束,但没有实现)

	// 1. 判断toUserId是不是ownerId的朋友，如果不是返回error
	fsi := FriendServiceImpl{}
	if t, err := fsi.IsFriend(toUserId, ownerId); err != nil {
		return err
	} else if !t {
		return ErrNotFriend
	}

	// 2. 将A -> B的消息插入数据库
	if err := mysql.CreateMessage(toUserId, ownerId, content); err != nil {
		return err
	}

	return nil
}
