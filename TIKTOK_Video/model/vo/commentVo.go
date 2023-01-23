package vo

import "TIKTOK_Video/model"

type CommentInfo struct {
	Id int64 `json:"id,omitempty"`
	// TODO:在完成http传输之后改为User
	User       int64  `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}
type Response struct {
	StatusCode int32         `json:"status_code"`
	StatusMsg  string        `json:"status_msg,omitempty"`
	Comment    model.Comment `json:"comment,omitempty"`
}
