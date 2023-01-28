package vo

type FollowResponse struct {
	Response
	UserInfoList []UserInfo `json:"user_list"`
}
