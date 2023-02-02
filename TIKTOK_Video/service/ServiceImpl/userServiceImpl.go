package ServiceImpl

import (
	"TIKTOK_Video/model/vo"
	"TIKTOK_Video/resolver"
	"context"
	"encoding/json"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	"strconv"
)

type UserServiceImpl struct {
}

var ErrGetUserInfo = errors.New("can not get the userInfo")

func (usi *UserServiceImpl) CreateUserByNameAndPassword(username, password string) (int64, error) {
	return 0, errors.New("CreateUserByNameAndPassword is un unsupported method")
}

func (usi *UserServiceImpl) GetUserInfoById(queryUserId int64, userId int64) (*vo.UserInfo, error) {
	//只查询一个id时也变成一个对象转到查询多个对象中
	mm, err := usi.GetUsersInfoByIds([]int64{queryUserId}, userId)
	if err != nil {
		return nil, err
	}
	if v, ok := mm[queryUserId]; ok {
		return v, nil
	}
	return nil, errors.New("wrong userId")
}

func (usi *UserServiceImpl) GetUsersInfoByIds(queryUserId []int64, userId int64) (map[int64]*vo.UserInfo, error) {
	//获得与注册中心同步的客户端对象
	cli := resolver.GetInstance()
	if cli == nil {
		return nil, ErrGetUserInfo
	}
	//发送http的参数
	args := &protocol.Args{}
	bytes, err := json.Marshal(queryUserId)
	if err != nil {
		return nil, err
	}
	args.Add("user_ids", string(bytes))
	args.Add("user_id", strconv.FormatInt(userId, 10))
	//发送请求，返回Response.StatusMsg为userinfo的数组json字符串
	status, body, err := cli.Post(context.Background(), nil, "http://tiktok.simple.user/douyin/user/get", args, config.WithSD(true))
	if status == 200 {
		res := vo.Response{}
		//userinfo反序列化成一个对象数组
		if err = json.Unmarshal(body, &res); err != nil {
			return nil, ErrGetUserInfo
		}
		users := make([]*vo.UserInfo, len(queryUserId))
		if err = json.Unmarshal([]byte(res.StatusMsg), &users); err != nil {
			return nil, ErrGetUserInfo
		}
		ret := make(map[int64]*vo.UserInfo)
		for _, user := range users {
			if user != nil {
				ret[user.Id] = user
			}

		}
		return ret, nil
	}
	return nil, ErrGetUserInfo
}
