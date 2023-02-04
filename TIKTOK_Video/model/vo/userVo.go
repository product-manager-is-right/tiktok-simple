package vo

/*
UserInfo 返回用户信息的实体类
*/
type UserInfo struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type UserInfosResponse struct {
	Response
	UserInfo []*UserInfo `json:"users,omitempty"`
}
