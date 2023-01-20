package mysql

/*
GetFollowCntByUserId
根据UserId查询该用户的关注数
*/
func GetFollowCntByUserId(userId int64) (int64, error) {
	// TODO : impl
	return 0, nil
}

/*
GetFollowerCntByUserId
根据UserId查询该用户的粉丝数/被关注数
*/
func GetFollowerCntByUserId(userId int64) (int64, error) {
	// TODO : impl
	return 0, nil
}

/*
GetIsFollow
判断userIdSrc 是否 关注 userIdDst
*/
func GetIsFollow(userIdDst, userIdSrc int64) (bool, error) {
	// TODO : impl
	return false, nil
}
