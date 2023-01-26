package ServiceImpl

import (
	"TIKTOK_Video/dal/mysql"
	"TIKTOK_Video/model"
	"TIKTOK_Video/model/vo"
	"errors"
	"math/rand"
	"strconv"
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
	for idx, comment := range comments {
		info, err := getUserInfoById(comment.UserId, userId)
		if err != nil {
			return nil, err
		}
		ret[idx] = &vo.CommentInfo{
			Id:         comment.Id,
			User:       *info,
			Content:    comment.Comment,
			CreateDate: time.Unix(comment.CreateDate, 0).Format("01-02"),
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
	userInfo, err := getUserInfoById(userId, userId)
	if err != nil {
		return nil, err
	}

	_, err = mysql.GetVideoByID(videoId)
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
		CreateDate: time.Unix(comment.CreateDate, 0).Format("01-02"),
	}
	return commentInfo, nil
}

// 调用远程接口，根据userid获取具体的user的个人信息
/**
userId：要查询的userId
ownerId: 发起查询评论的ID，用于判断是否是followed
@return map 键为userId，值为用户信息的map
@return error 错误信息，没有错误就返回nil

*/
func getUserInfoById(userId int64, ownerId int64) (*vo.UserInfo, error) {
	//TODO 远程调用获取User信息
	//先手动实现一下
	rand.Seed(time.Now().Unix())
	ret := &vo.UserInfo{
		Id:            userId,
		Name:          "user" + strconv.FormatInt(userId, 10),
		FollowCount:   rand.Int63() % 10000,
		FollowerCount: rand.Int63() % 10000,
		IsFollow:      userId%73 == 0,
	}
	return ret, nil
	// TODO : impl
	//return nil, nil
}

// 调用远程接口，根据userid数组获取具体的user的个人信息
/**
userIds：要查询的userIds
ownerId: 发起查询评论的ID，用于判断是否是followed
@return map 键为userId，值为用户信息的map
@return error 错误信息，没有错误就返回nil

*/
func getUserInfoByIds(userIds []int64, ownerId int64) (map[int64]*vo.UserInfo, error) {
	//先手动实现一下
	mm := make(map[int64]*vo.UserInfo, len(userIds))
	rand.Seed(time.Now().Unix())
	for _, userId := range userIds {
		mm[userId] = &vo.UserInfo{
			Id:            userId,
			Name:          "user" + strconv.FormatInt(userId, 10),
			FollowCount:   rand.Int63() % 10000,
			FollowerCount: rand.Int63() % 10000,
			IsFollow:      userId%73 == 0,
		}
	}
	return mm, nil
	// TODO : impl
	//return nil, nil
}
