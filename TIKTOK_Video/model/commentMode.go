package model

// 评论信息
type Comment struct {
	Id int64
	// TODO:在完成http传输之后改为User
	User       int64
	Content    string
	CreateDate string
}
