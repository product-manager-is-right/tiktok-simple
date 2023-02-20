package mysql

import (
	"TIKTOK_Video/model"
	"errors"
	"gorm.io/gorm"
	"time"
)

func GetCommentByID(id int64) (*model.Comment, error) {
	res := &model.Comment{}
	result := DB.Where("id = ?", id).
		Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}

	return res, nil
}

func GetCommentByVideoIds(videoId int64) ([]*model.Comment, error) {
	obj := model.Comment{}
	ret := make([]*model.Comment, 0)
	DB.Model(&obj).Where("video_id = ?", videoId).Order("create_date desc").Find(&ret)
	return ret, nil
}

func InsertComment(commentText string, videoId, userId int64) (*model.Comment, error) {
	//封装评论
	comment := model.Comment{
		Id:         0,
		VideoId:    videoId,
		UserId:     userId,
		Comment:    commentText,
		CreateDate: time.Now().Unix(),
	}
	result := DB.Create(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("新增失败")
	}
	return &comment, nil
}

func DeleteCommentByCommentId(commentId int64) error {
	comment := model.Comment{
		Id: commentId,
	}
	result := DB.Delete(&comment)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("删除失败")
	}
	return nil
}

func DecrementCommentCount(videoId int64) error {
	video := model.Video{VideoId: videoId}
	result := DB.Model(&video).Where("video_id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - 1"))
	var err error
	if err = result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return errors.New("减少评论数失败")
	}
	return nil
}

func IncrementCommentCount(videoId int64) error {
	video := model.Video{VideoId: videoId}
	result := DB.Model(&video).Where("video_id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + 1"))
	var err error
	if err = result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return errors.New("添加评论数失败")
	}
	return nil
}
