package api

import (
	"errors"

	"github.com/lgyong511/sffld-go/dao/user"
	"github.com/lgyong511/sffld-go/model"
	"github.com/lgyong511/sffld-go/util/jwt"
)

// Login 登录
func Login(u model.User) (string, error) {
	// 从db获取用户信息
	user, err := user.GetUserByName(u.Username)
	if err != nil {
		return "", err
	}
	// 校验用户
	if user.Username != u.Username && user.Password != u.Password {
		return "", errors.New("username or password error")
	}

	// 生成token
	token, err := jwt.GenToken(u.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

// SignUp 注册
func Register(u *model.User) error {
	// 从db获取用户信息
	_, err := user.GetUserByName(u.Username)
	if err == nil {
		return errors.New("username already exists")
	}
	// 注册用户
	return user.CreateUser(u)
}

// Logout 注销
func Logout(username string) error {
	// 从db删除token
	return nil
}
