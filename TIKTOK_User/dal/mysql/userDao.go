package mysql

import (
	"TIKTOK_User/model"
	"TIKTOK_User/util"
	"errors"
)

// CreateUser
// 创建用户
func CreateUser(username string, password string) (int64, error) {
	user := model.User{Username: username, Password: util.MD5(password)}
	result := DB.Create(&user)
	return user.Id, result.Error
}

// GetUserByUserName
// 通过用户名查找User
func GetUserByUserName(username string) (*model.User, error) {
	var res *model.User
	if err := DB.Where("username = ?", username).
		First(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CheckUser
// 检验用户名和密码是否正确
func CheckUser(username string, password string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := DB.Where("username = ? AND password = ?", username, password).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetUserByUserId
// 通过UserId获取User
func GetUserByUserId(userId int64) (*model.User, error) {
	res := &model.User{}
	if err := DB.Where("id = ?", userId).
		First(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func GetUserByIds(userIds []int64) ([]*model.User, error) {
	users := make([]*model.User, 0)
	result := DB.Find(&users, userIds)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.New("find error")
	}
	return users, nil

}
