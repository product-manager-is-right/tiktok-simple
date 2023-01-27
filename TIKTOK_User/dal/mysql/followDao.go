package mysql

import (
	"GoProject/model"
	"log"
)

/*
GetFollowCntByUserId
根据UserId查询该用户的关注数
*/
func GetFollowCntByUserId(userId int64) (int64, error) {
	// TODO : impl
	//return 1, nil
	var cnt int64 = 10
	// 当查询出现错误的情况，日志打印err msg，并返回err.
	if err := DB.
		Model(model.Follow{}).
		Where("user_id_from = ?", userId).
		Where("cancel = ?", 0).
		Count(&cnt).Error; nil != err {
		log.Println(err.Error())
		return 0, err
	}
	// 正常情况，返回取到的粉丝数。
	log.Printf("有%d个粉丝", cnt)
	return cnt, nil
}

/*
GetFollowerCntByUserId
根据UserId查询该用户的粉丝数/被关注数
*/
func GetFollowerCntByUserId(userId int64) (int64, error) {
	// TODO : impl
	// 用于存储当前用户粉丝数的变量
	var cnt int64
	// 当查询出现错误的情况，日志打印err msg，并返回err.
	if err := DB.
		Model(model.Follow{}).
		Where("user_id_to = ?", userId).
		Where("cancel = ?", 0).
		Count(&cnt).Error; nil != err {
		log.Println(err.Error())
		log.Printf("有%d个关注", cnt)
		return 0, err
	}
	// 正常情况，返回取到的粉丝数。
	return cnt, nil
}

/*
GetIsFollow
判断userIdSrc 是否 关注 userIdDst
*/
func GetIsFollow(userIdDst, userIdSrc int64) (bool, error) {
	// TODO : impl
	// 用于存储查出来的关注关系。
	follow := model.Follow{}

	//notIsNull := false
	//当查询出现错误时，日志打印err msg，并return err.
	if err := DB.
		Where("user_id_from = ?", userIdSrc).
		Where("user_id_to = ?", userIdDst).
		Where("cancel = ? or cancel = ?", 0, 1).
		Take(&follow).Error; nil != err {
		// 当没查到记录报错时，不当做错误处理。
		if "record not found" == err.Error() {
			return false, nil //return nil,nil  只能传参为(*Follow, error)

		}
		log.Println(err.Error())
		return false, err //return nil,err 只能传参为([]int64, error)
	}
	//正常情况，返回取到的关系和空err.
	//return &follow, nil  r如果是查询对应的关注关系，直接返回 &follow，前提是传参为([]int64, error)
	return true, nil
}
