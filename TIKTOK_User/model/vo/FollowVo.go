package vo

type FollowResponse struct {
	Response
	UserInfoList []UserInfo `json:"user_list"`
}

type FollowerResponse struct {
	Response
	UserInfoList []UserInfo `json:"user_list"`
}
type FollowActionResponse struct {
	Response
}

type RelationResponse struct {
	Response
	FriendUserInfoList []FriendUserInfo `json:"user_list"`
}
