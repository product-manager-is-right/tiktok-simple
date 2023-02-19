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
	WorkCount     int64  `json:"work_count"`
	FavoriteCount int64  `json:"favorite_count"`
	Avatar        string `json:"avatar"`           //头像，写死
	Signature     string `json:"signature"`        //签名
	Background    string `json:"background_image"` //背景图片写死
}

type UserInfosResponse struct {
	Response
	UserInfo []*UserInfo `json:"users,omitempty"`
}
