package mysql

import (
	"TIKTOK_User/model"
	"github.com/go-errors/errors"
	"gorm.io/gorm"
	"log"
)

type FollowDao struct {
}

/*
GetFollowCntByUserId
根据UserId查询该用户的关注数
*/
func GetFollowCntByUserId(userId int64) (int64, error) {
	var cnt int64 = 0

	if err := DB.Model(model.Follow{}).Where("user_id_from = ?", userId).
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

	if err := DB.Model(model.Follow{}).
		Where("user_id_to = ?", userId).
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
func GetIsFollow(userTo, userFrom int64) (bool, error) {
	follow := model.Follow{}

	if err := DB.Where("user_id_from = ?", userFrom).
		Where("user_id_to = ?", userTo).
		Take(&follow).Error; err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil

}

func CreateRelation(userToId, userFromId int64) error {
	Follow := model.Follow{UserIdTo: userToId, UserIdFrom: userFromId}
	result := DB.Create(&Follow)
	return result.Error
}

func GetRelation(userToId, userFromId int64) (*model.Follow, error) {
	var res *model.Follow
	if err := DB.Where("user_id_to = ?", userToId).Where("user_id_from = ?", userFromId).
		First(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteRelation(userToId, userFromId int64) error {
	follow := &model.Follow{}
	if err := DB.Model(follow).Where("user_id_to = ?", userToId).Where("user_id_from = ?", userFromId).
		Delete(follow).Error; err != nil {
		return errors.New("关系删除失败")
	}
	return nil
}

/*
GetFollowingIds
给定用户id，查询他关注了哪些人的id。
*/
func GetFollowingIds(userId int64) ([]int64, error) {
	var ids []int64
	if err := DB.Model(model.Follow{}).Where("user_id_from = ?", userId).
		Pluck("user_id_to", &ids).Error; err != nil {
		return nil, err
	}
	// 查询成功。
	return ids, nil
}

// GetFollowerIds 给定用户id，查询他的粉丝id列表
func GetFollowerIds(userId int64) ([]int64, error) {
	var ids []int64
	if err := DB.Model(model.Follow{}).Where("user_id_to = ?", userId).
		Pluck("user_id_from", &ids).Error; err != nil {
		return nil, err
	}
	// 查询成功。
	return ids, nil
}
