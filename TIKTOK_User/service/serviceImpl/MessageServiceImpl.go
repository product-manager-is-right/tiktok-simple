package serviceImpl

import (
	"TIKTOK_User/dal/mysql"
	"TIKTOK_User/model/vo"
	"errors"
	"sort"
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

func (msi *MessageServiceImpl) GetMessage(toUserId int64, ownerId int64, preTime int64) ([]vo.MessageInfo, error) {
	var res []vo.MessageInfo
	//检查最新消息的时间
	m1 := mysql.GetLatestMessage(toUserId, ownerId)
	m2 := mysql.GetLatestMessage(ownerId, toUserId)
	if m1.CreateTime <= preTime || m2.CreateTime <= preTime {
		return res, nil
	}
	messageList, err := mysql.GetMessage(toUserId, ownerId)
	if err != nil {
		return res, err
	}
	for _, message := range messageList {
		if message.CreateTime > preTime {
			m := vo.MessageInfo{
				ID:         message.Id,
				ToUserId:   message.UserIdTo,
				FromUserId: message.UserIdFrom,
				Content:    message.Message,
				CreateTime: message.CreateTime,
			}
			res = append(res, m)
		}
	}
	sort.Sort(Messages(res))
	return res, nil
}

// GetLatestMessage 获取fromUser和toUser的最新的一条聊天记录
// @param : message 消息
// @param : msgType 消息类型 0 - toUser发送 1 - fromUser发送
func (msi *MessageServiceImpl) GetLatestMessage(toUserId, fromUserId int64) (message string, msgType int64) {
	m := mysql.GetLatestMessage(toUserId, fromUserId)
	if m.UserIdFrom == fromUserId {
		return m.Message, 1
	}
	return m.Message, 0
}

// 结构体数组排序
type Messages []vo.MessageInfo

// 排序函数实现
func (m Messages) Len() int {
	return len(m)
}
func (m Messages) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m Messages) Less(i, j int) bool {
	return m[i].CreateTime < m[j].CreateTime
}
