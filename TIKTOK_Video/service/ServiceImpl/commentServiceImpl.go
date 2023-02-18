package ServiceImpl

import (
	"TIKTOK_Video/dal/mysql"
	"TIKTOK_Video/model"
	"TIKTOK_Video/model/vo"
	"TIKTOK_Video/mw/rabbitMQ/producer"
	"TIKTOK_Video/mw/redis"
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
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
	var commentInfos []*vo.CommentInfo
	// 先从redis中查找
	strVideoId := strconv.FormatInt(videoId, 10)
	vs, err := redis.CommentList.Get(context.Background(), strVideoId).Result()
	// 缓存命中
	if err == nil {
		// 反序列化
		if err := json.Unmarshal([]byte(vs), &commentInfos); err != nil {
			return nil, errors.New("redis Unmarshal:" + err.Error())
		}
		return commentInfos, nil
	}
	//根据videoId获取comment
	comments, err := mysql.GetCommentByVideoIds(videoId)
	if err != nil {
		return nil, err
	}
	commentInfos, err = bindCommentInfo(comments, userId)
	if err != nil {
		return nil, err
	}
	//序列化存入redis
	strCommentInfos, _ := json.Marshal(commentInfos)
	redis.CommentList.Set(context.Background(), strVideoId, strCommentInfos, redis.SetExpiredTime())
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
	// 判断评论是否存在以及是否有权限访问
	c, err := mysql.GetCommentByID(commentId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("该评论不存在")
		}
		return err
	}
	if c.UserId != userId {
		return errors.New("无权限操作该评论")
	}

	//首先尝试发送处理到mq中
	if err := producer.SendDelCommentMessage(commentId, videoId); err != nil {
		//发送失败。自动同步操作数据库.  这两条应该是一个事务。
		if err = mysql.DeleteCommentByCommentId(commentId); err != nil {
			return err
		}
		if err = mysql.DecrementCommentCount(videoId); err != nil {
			return err
		}
		strVideoId := strconv.FormatInt(videoId, 10)
		for i := 0; i < redis.RetryTime; i++ {
			if _, err := redis.CommentList.Del(context.Background(), strVideoId).Result(); err == nil {
				break
			}
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
	strVideoId := strconv.FormatInt(videoId, 10)
	for i := 0; i < redis.RetryTime; i++ {
		if _, err := redis.CommentList.Del(context.Background(), strVideoId).Result(); err == nil {
			break
		}
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
