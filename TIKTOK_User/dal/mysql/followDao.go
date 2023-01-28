package mysql

import (
	"GoProject/model"
	"gorm.io/gorm"
	"log"
)

/*
GetFollowCntByUserId
根据UserId查询该用户的关注数
*/
func GetFollowCntByUserId(userId int64) (int64, error) {
	var cnt int64 = 10

	if err := DB.Model(model.Follow{}).Where("user_id_from = ?", userId).Where("cancel = ?", 0).
		Count(&cnt).Error; err != nil {
		log.Println(err.Error())
		return 0, err
	}

	return cnt, nil
}

/*
GetFollowerCntByUserId
根据UserId查询该用户的粉丝数/被关注数
*/
func GetFollowerCntByUserId(userId int64) (int64, error) {
	var cnt int64

	if err := DB.Model(model.Follow{}).Where("user_id_to = ?", userId).Where("cancel = ?", 0).
		Count(&cnt).Error; err != nil {
		log.Println(err.Error())
		return 0, err
	}

	return cnt, nil
}

/*
GetIsFollow
判断userIdSrc 是否 关注 userIdDst
*/
func GetIsFollow(userIdDst, userIdSrc int64) (bool, error) {
	follow := model.Follow{}

	if err := DB.Where("user_id_from = ?", userIdSrc).
		Where("user_id_to = ?", userIdDst).
		Where("cancel = ?", 0).
		Take(&follow).Error; err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}
