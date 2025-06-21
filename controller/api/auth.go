package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lgyong511/sffld-go/controller"
	"github.com/lgyong511/sffld-go/model"
	"github.com/lgyong511/sffld-go/util/jwt"
	"github.com/sirupsen/logrus"
)

// Login 登录
func Login(c *gin.Context) {
	//解析参数
	user := new(model.User)
	if err := c.ShouldBind(user); err != nil {
		logrus.Error(err)
		controller.Error(c, controller.ParamsErrorCode, controller.CodeToMsg(controller.ParamsErrorCode))
		return
	}

	// 校验参数
	if user.Username == "" || user.Password == "" {
		controller.Error(c, controller.ParamsErrorCode, controller.CodeToMsg(controller.ParamsErrorCode))
		return
	}

	// 校验用户
	if user.Username != "admin" || user.Password != "admin" {
		logrus.Error("username or password error")
		controller.Error(c, controller.UsernameOrPasswordErrorCode, controller.CodeToMsg(controller.UsernameOrPasswordErrorCode))
		return
	}

	// 生成token
	token, err := jwt.GenToken(user.Username)
	if err != nil {
		logrus.Error("gen token error")
		controller.Error(c, controller.GenTokenErrorCode, controller.CodeToMsg(controller.GenTokenErrorCode))
		return
	}

	// 返回token
	controller.Token(c, token)
}
