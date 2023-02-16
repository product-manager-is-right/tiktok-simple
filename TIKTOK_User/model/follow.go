package model

type Follow struct {
	Id         int64 `column:"id"`
	UserIdFrom int64 `column:"user_id_from"`
	UserIdTo   int64 `column:"user_id_to"`
}

func (f *Follow) TableName() string {
	return "ums_follow"
}
