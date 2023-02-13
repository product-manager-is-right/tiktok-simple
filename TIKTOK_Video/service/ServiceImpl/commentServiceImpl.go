package ServiceImpl

import (
	"TIKTOK_Video/dal/mysql"
	"TIKTOK_Video/model"
	"TIKTOK_Video/model/vo"
	"TIKTOK_Video/mw/rabbitMQ"
	"errors"
	"log"
	"strconv"
	"strings"
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

func (csi *CommentServiceImpl) DeleteCommentByCommentId(commentId, userId, videoId int64) error {
	if commentId <= 0 {
		return errors.New("wrong commentId")
	}
	//获取评论的具体信息，如果没有记录会返回err。2.13被阿耿删减，因为使用了mq进行异步处理了
	//这一步好像没有什么用处了，如果没有这个评论或者videoId对不上的话在mq的处理中直接无视就好了
	//comment, err := mysql.GetCommentByID(commentId)
	//if err != nil {
	//	return err
	//}

	//首先尝试发送处理到mq中
	if err := sendDelMessage(commentId, userId, videoId); err != nil {
		//发送失败。自动同步操作数据库
		if err = mysql.DeleteCommentByCommentId(commentId, userId); err != nil {
			return err
		}
		if err = mysql.DecrementCommentCount(videoId); err != nil {
			return err
		}
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

// 发送顺序为commentId, userId, videoId
func sendDelMessage(id ...int64) error {
	//使用最高的36，压缩一下
	if id == nil || len(id) <= 0 {
		return errors.New("cannot send empty message to rabbitmq")
	}
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatInt(id[0], 36))
	for i := 1; i < len(id); i++ {
		sb.WriteString("-")
		sb.WriteString(strconv.FormatInt(id[i], 36))
	}
	if err := rabbitMQ.RmqComment.Publish(sb.String()); err != nil {
		log.Print(err)
		return err
	}
	return nil
}
