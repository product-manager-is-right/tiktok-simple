package service

import (
	"TIKTOK_Video/model/vo"
	"TIKTOK_Video/service/ServiceImpl"
	"sync"
)

type CommentService interface {
	// GetCommentListByVideoId 通过视频ID查询评论列表指针
	GetCommentListByVideoId(videoId, userId int64) (commentList []*vo.CommentInfo, err error)
	GetCommentByCommentId(commentId, userId int64) (comment *vo.CommentInfo, err error)
	// InsertComment 插入新的评论内容，成功返回CommentInfo的指针
	InsertComment(commentText string, videoId, userId int64) (comment *vo.CommentInfo, err error)
	DeleteCommentByCommentId(commentId, userId int64) error
	//ReturnA() string
}

var (
	service            CommentService
	CommentServiceOnce sync.Once
)

// NewCommentServiceInstance  单例模式返回service对象
func NewCommentServiceInstance() CommentService {
	CommentServiceOnce.Do(
		func() {
			service = &ServiceImpl.CommentServiceImpl{}
		})
	return service
}
