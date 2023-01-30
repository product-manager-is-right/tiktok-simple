package vo

type FollowResponse struct {
	Response
	UserInfoList []UserInfo `json:"user_list"`
}

type FollowerResponse struct {
	Response
	UserInfoList []UserInfo `json:"user_list"`
}

type RelationResponse struct {
	Response
	UserInfoList []UserInfo `json:"user_list"`
}
