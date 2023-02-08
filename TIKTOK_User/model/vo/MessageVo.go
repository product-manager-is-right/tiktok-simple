package vo

type MessageInfo struct {
	ID         int64  `json:"id"`          // 消息id
	Content    string `json:"content"`     // 消息内容
	CreateTime string `json:"create_time"` // 消息发送时间 yyyy-MM-dd HH:MM:ss
}

type MessageActionResponse struct {
	Response
}

type ChatResponse struct {
	Response
	MessageList []MessageInfo `json:"message_list"`
}
