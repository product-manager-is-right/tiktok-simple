package vo

/*
CommentInfo 返回评论信息的实体类
*/
type CommentInfo struct {
	Id         int64    `json:"id"`
	User       UserInfo `json:"user"`
	Content    string   `json:"content"`
	CreateDate string   `json:"create_date"`
}

/*
CommentResponse
单个评论的响应
*/
type CommentResponse struct {
	Response
	Comment CommentInfo `json:"comment"`
}

/*
CommentListResponse
评论列表响应实体
*/
type CommentListResponse struct {
	Response
	//用指针好像就不用结构体拷贝，会快一点
	CommentList []*CommentInfo `json:"comment_list"`
	//CommentList []CommentInfo  `json:"comment_list"`
}
