package vo

type MessageInfo struct {
	ID         int64  `json:"id"`           // 消息id
	ToUserId   int64  `json:"to_user_id"`   // 该消息接收者的id
	FromUserId int64  `json:"from_user_id"` // 该消息发送者的id
	Content    string `json:"content"`      // 消息内容
	CreateTime string `json:"create_time"`  // 消息发送时间 yyyy-MM-dd HH:MM:ss
}

type MessageActionResponse struct {
	Response
}

type ChatResponse struct {
	Response
	MessageList []MessageInfo `json:"message_list"`
}
