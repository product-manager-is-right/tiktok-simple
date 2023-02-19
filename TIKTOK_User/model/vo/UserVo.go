package vo

/*
UserInfo 返回用户信息的实体类
*/
//TODO 需要修添加用户发布视频数量
type UserInfo struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	WorkCount     int64  `json:"work_count"`
	FavoriteCount int64  `json:"favorite_count"`
	Avatar        string `json:"avatar"`           //头像，写死
	Signature     string `json:"signature"`        //签名
	Background    string `json:"background_image"` //背景图片写死
}

type FriendUserInfo struct {
	UserInfo
	Message string `json:"message"`
	MsgType int64  `json:"msgType"` // 0接受 or 1发送
}

/*
Response Vo
*/
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserInfoResponse struct {
	Response
	UserInfo *UserInfo `json:"user,omitempty"`
}

// UserInfosResponse 注意和上面UserInfoResponse的区别，这里是返回UserInfo数组
type UserInfosResponse struct {
	Response
	UserInfo []*UserInfo `json:"users,omitempty"`
}

type RegisterResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type LoginResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
