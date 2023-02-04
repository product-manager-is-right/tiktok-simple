package ServiceImpl

import (
	"TIKTOK_Video/dal/mysql"
	"TIKTOK_Video/model"
	"TIKTOK_Video/model/vo"
	"errors"
	"time"
)

type CommentServiceImpl struct {
}

// GetCommentListByVideoId 根据视频ID获取评论列表
func (csi *CommentServiceImpl) GetCommentListByVideoId(videoId, userId int64) ([]*vo.CommentInfo, error) {
	if videoId < 0 {
		return nil, errors.New("wrong videoId")
	}
	//根据videoId获取comment
	comments, err := mysql.GetCommentByVideoIds(videoId)
	if err != nil {
		return nil, err
	}
	commentInfos, err := bindCommentInfo(comments, userId)
	if err != nil {
		return nil, err
	}
	return commentInfos, nil
}

func bindCommentInfo(comments []*model.Comment, userId int64) ([]*vo.CommentInfo, error) {
	ret := make([]*vo.CommentInfo, len(comments))
	//将需要查询的id拿出来
	userIds := make([]int64, len(comments))
	for i, comment := range comments {
		userIds[i] = comment.UserId
	}
	serviceImpl := UserServiceImpl{}
	userInfo, err := serviceImpl.GetUsersInfoByIds(userIds, userId)
	if err != nil {
		return nil, err
	}
	for idx, comment := range comments {
		ret[idx] = &vo.CommentInfo{
			Id:         comment.Id,
			User:       *userInfo[comment.UserId],
			Content:    comment.Comment,
			CreateDate: time.Unix(comment.CreateDate, 0).Format("01-02 15:04"),
		}
	}
	return ret, nil
}

func (csi *CommentServiceImpl) GetCommentByCommentId(commentId, userId int64) (commentInfo *vo.CommentInfo, err error) {
	if commentId <= 0 {
		return nil, errors.New("wrong commentId")
	}
	comment, err := mysql.GetCommentByID(commentId)
	if err != nil {
		return nil, err
	}
	arr := make([]*model.Comment, 1)
	arr[0] = comment
	commentInfos, err := bindCommentInfo(arr, userId)
	if err != nil {
		return nil, err
	}
	return commentInfos[0], nil

}

func (csi *CommentServiceImpl) ReturnA() string {
	return "a"
}

func (csi *CommentServiceImpl) DeleteCommentByCommentId(commentId, userId int64) error {
	if commentId <= 0 {
		return errors.New("wrong commentId")
	}
	comment, err := mysql.GetCommentByID(commentId)
	if err != nil {
		return err
	}
	if err = mysql.DeleteCommentByCommentId(commentId, userId); err != nil {
		return err
	}
	err = mysql.DecrementCommentCount(comment.VideoId)
	if err != nil {
		return err
	}
	return nil

}

func (csi *CommentServiceImpl) InsertComment(commentText string, videoId, userId int64) (commentInfo *vo.CommentInfo, err error) {
	//由于没有外键约束，手动检查是否存在这个videoId和是不是userID对得上
	_, err = mysql.GetVideoByID(videoId)
	if err != nil {
		return nil, err
	}
	serviceImpl := UserServiceImpl{}
	userInfo, err := serviceImpl.GetUserInfoById(userId, userId)
	if err != nil {
		return nil, err
	}
	comment, err := mysql.InsertComment(commentText, videoId, userId)
	if err != nil {
		return nil, err
	}
	//到这里没问题，给video的comment_count字段加一
	if err = mysql.IncrementCommentCount(videoId); err != nil {
		return nil, err
	}
	//封装commentInfo对象
	commentInfo = &vo.CommentInfo{
		Id:         comment.Id,
		User:       *userInfo,
		Content:    comment.Comment,
		CreateDate: time.Unix(comment.CreateDate, 0).Format("01-02 15:04"),
	}
	return commentInfo, nil
}
