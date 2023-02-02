package mysql

import (
	"TIKTOK_User/model"
	"errors"
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
	var cnt int64 = 10

	if err := DB.Model(model.Follow{}).Where("user_id_from = ?", userId).Where("cancel = ?", 0).
		Count(&cnt).Error; err != nil {
		log.Println(err.Error())
		return 0, err
	}

	return cnt, nil
	// TODO : impl
	//return 1, nil
	var cnt int64 = 10

	if err := DB.Model(model.Follow{}).
		Where("user_id_from = ?", userId).
		Where("cancel = ?", 0).
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
	// TODO : impl
	//return 1, nil
	var cnt int64

	if err := DB.Model(model.Follow{}).
		Where("user_id_to = ?", userId).
		Where("cancel = ?", 0).
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
func GetIsFollow(usertoid, userfromid int64) (bool, error) {
	// TODO : impl
	//return false, nil
	follow := model.Follow{}

	if err := DB.Where("user_id_from = ?", userfromid).
		Where("user_id_to = ?", usertoid).
		Where("cancel = ?", 0).
		Take(&follow).Error; err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil

}

func CreateNewrelation(usertoid, userfromid, Cancel int64) (int64, error) {
	Follow := model.Follow{UserIdTo: usertoid, UserIdFrom: userfromid, Cancel: 0}
	result := DB.Create(&Follow)
	return Follow.Id, result.Error
}

func Getrelation(usertoid, userfromid int64) ([]*model.Follow, error) {
	res := make([]*model.Follow, 0)
	if err := DB.Where("user_id_to = ?", usertoid).Where("user_id_from = ?", userfromid).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func Deleterelation(usertoid, userfromid int64) error {
	Follow := model.Follow{UserIdTo: usertoid, UserIdFrom: userfromid, Cancel: 0}

	result := DB.Where("user_id_from = ?", userfromid).Where("user_id_to = ?", usertoid).
		Delete(&Follow)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("删除失败")
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
		Where("cancel = ?", 0).Pluck("user_id_to", &ids).Error; nil != err {
		if "record not found" == err.Error() {
			return nil, nil
		}
		// 查询出错。
		log.Println(err.Error())
		return nil, err
	}
	// 查询成功。
	return ids, nil
}

func GetFollowerIds(userId int64) ([]int64, error) {
	var ids []int64
	if err := DB.Model(model.Follower{}).Where("user_id_from = ?", userId).
		Where("cancel = ?", 0).Pluck("user_id_to", &ids).Error; nil != err {
		if "record not found" == err.Error() {
			return nil, nil
		}
		// 查询出错。
		log.Println(err.Error())
		return nil, err
	}
	return ids, nil
}

/*
GetFriendsIds
给定用户id，查询他好友的id。
*/
func GetFriendsIds(userId int64) ([]int64, error) {
	var ids []int64
	if err := DB.Model(model.Follow{}).Where("user_id_from = ?", userId).
		Where("cancel = ?", 0).Pluck("user_id_to", &ids).Error; nil != err {
		if "record not found" == err.Error() {
			return nil, nil
		}
		// 查询出错。
		log.Println(err.Error())
		return nil, err
	}
	// 查询成功。
	return ids, nil
}
