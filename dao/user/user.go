package user

import (
	"github.com/lgyong511/sffld-go/model"
)

// 判断用户名是否存在
func IsUsernameExist(username string) bool {
	return username == "admin"
}

// GetUserByName 根据用户名获取用户
func GetUserByName(username string) (*model.User, error) {

	return &model.User{
		Username: username,
		Password: "admin",
	}, nil
}

// CreateUser 创建用户
func CreateUser(u *model.User) error {
	return nil
}
