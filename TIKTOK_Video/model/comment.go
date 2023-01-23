package model

// Comment 评论信息
type Comment struct {
	Id         int64
	VideoId    int64
	UserId     int64
	Comment    string
	CreateDate int64
}

func (c *Comment) TableName() string {
	return "vms_comment"
}
