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

/*
VideoInfo 返回视频信息的实体类
*/
type VideoInfo struct {
	Id            int64    `json:"id,omitempty"`
	Author        UserInfo `json:"author"`
	PlayUrl       string   `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`
	FavoriteCount int64    `json:"favorite_count,omitempty"`
	CommentCount  int64    `json:"comment_count,omitempty"`
	IsFavorite    bool     `json:"is_favorite,omitempty"`
	Title         string   `json:"title,omitempty"`
}

/*
Response Vo
*/
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type FeedResponse struct {
	Response
	VideoList []VideoInfo `json:"video_list"`
	NextTime  int64       `json:"next_time"`
}
type VideoInfoResponse struct {
	Response
	VideoList []VideoInfo `json:"video_list,omitempty"`
}
