package mysql

import (
	"TIKTOK_Video/model"
	"errors"
	"time"
)

func GetCommentByID(id int64) (*model.Comment, error) {
	res := &model.Comment{}
	result := DB.Where("id = ?", id).
		Find(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("查找失败")
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

func DeleteCommentByCommentId(commentId, userId int64) error {
	comment := model.Comment{
		Id:         commentId,
		VideoId:    0,
		UserId:     userId,
		Comment:    "",
		CreateDate: 0,
	}
	result := DB.Where("user_id = ?", userId).Delete(&comment)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("删除失败")
	}
	return nil
}
