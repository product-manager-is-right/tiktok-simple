package model

type Message struct {
	Id         int64  `column:"id"`
	UserIdFrom int64  `column:"user_id_from"`
	UserIdTo   int64  `column:"user_id_to"`
	Message    string `column:"message"`
	CreateTime int64  `column:"create_time"`
}

func (m *Message) TableName() string {
	return "ums_message"
}
