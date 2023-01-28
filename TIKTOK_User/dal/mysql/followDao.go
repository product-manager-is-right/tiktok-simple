package mysql

import (
	"GoProject/model"
	"log"
)

type FollowDao struct {
}

/*
GetFollowCntByUserId
根据UserId查询该用户的关注数
*/
func GetFollowCntByUserId(userId int64) (int64, error) {
	// TODO : impl
	return 1, nil
}

/*
GetFollowerCntByUserId
根据UserId查询该用户的粉丝数/被关注数
*/
func GetFollowerCntByUserId(userId int64) (int64, error) {
	// TODO : impl
	return 1, nil
}

/*
GetIsFollow
判断userIdSrc 是否 关注 userIdDst
*/
func GetIsFollow(userIdDst, userIdSrc int64) (bool, error) {
	// TODO : impl
	return false, nil
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
