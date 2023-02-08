package serviceImpl

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
	"errors"
	"time"
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

func (msi *MessageServiceImpl) GetMessage(toUserId int64, ownerId int64) ([]vo.MessageInfo, error) {
	var res []vo.MessageInfo

	messageList, err := mysql.GetMessage(toUserId, ownerId)
	if err != nil {
		return res, err
	}
	for _, message := range messageList {
		var cur vo.MessageInfo
		cur.ID = message.Id
		//cur.CreateTime = strconv.FormatInt(message.CreateTime, 10)
		cur.CreateTime = unixToStr(message.CreateTime, "2006-01-02 15:04:05")
		cur.Content = message.Message
		res = append(res, cur)
	}
	return res, nil
}

// 时间戳转时间
func unixToStr(timeUnix int64, layout string) string {
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	return timeStr
}
