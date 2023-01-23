package mysql

import "TIKTOK_Video/model"

func GetCommentByID(id int64) (*model.Comment, error) {
	res := &model.Comment{}
	if err := DB.Where("id = ?", id).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
