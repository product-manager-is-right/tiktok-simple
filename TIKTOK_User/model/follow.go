package model

type Follow struct {
	Id         int64 `column:"id"`
	UserIdFrom int64 `column:"user_id_from"`
	UserIdTo   int64 `column:"user_id_to"`
	Cancel     int64 `column:"cancel"`
}

func (f *Follow) TableName() string {
	return "ums_follow"
}

type Follower struct {
	Id         int64 `column:"id"`
	UserIdFrom int64 `column:"user_id_from"`
	UserIdTo   int64 `column:"user_id_to"`
	Cancel     int64 `column:"cancel"`
}

func (f *Follower) TableName() string {
	return "ums_friend"
}
