package vo

/*
UserInfo 返回用户信息的实体类
*/
type UserInfo struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

/*
VideoInfo 返回视频信息的实体类
*/
type VideoInfo struct {
	Id            int64    `json:"id"`
	Author        UserInfo `json:"user"`
	PlayUrl       string   `json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int64    `json:"favorite_count"`
	CommentCount  int64    `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   `json:"title"`
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
