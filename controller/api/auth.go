package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lgyong511/sffld-go/controller"
	"github.com/lgyong511/sffld-go/model"
	"github.com/lgyong511/sffld-go/service/api"
	"github.com/lgyong511/sffld-go/util/jwt"
	"github.com/sirupsen/logrus"
)

// Login 登录handler
func Login(c *gin.Context) {
	//解析参数
	user := new(model.User)
	if err := c.ShouldBind(user); err != nil {
		logrus.Error(err)
		controller.Error(c, controller.ParamsErrorCode, controller.CodeToMsg(controller.ParamsErrorCode))
		return
	}

	// 登录
	token, err := api.Login(*user)
	if err != nil {
		logrus.Error(err)
		controller.Error(c, controller.ParamsErrorCode, controller.CodeToMsg(controller.ParamsErrorCode))
		return
	}

	// 返回token
	controller.Login(c, token)
}

// Register 注册handler
func Register(c *gin.Context) {
	//解析参数
	user := new(model.User)
	if err := c.ShouldBind(user); err != nil {
		logrus.Error(err)
		controller.Error(c, controller.ParamsErrorCode, controller.CodeToMsg(controller.ParamsErrorCode))
		return
	}

	// 注册
	err := api.Register(user)
	if err != nil {
		logrus.Error(err)
		controller.Error(c, controller.ParamsErrorCode, controller.CodeToMsg(controller.ParamsErrorCode))
		return
	}
	controller.Success(c, nil)
}

// Logout 注销handler
func Logout(c *gin.Context) {
	claims, _ := c.Get("claims")

	// 注销
	err := api.Logout(claims.(*jwt.CustomClaims).Username)
	if err != nil {
		logrus.Error(err)
		controller.Error(c, controller.ParamsErrorCode, controller.CodeToMsg(controller.ParamsErrorCode))
		return
	}
	controller.Success(c, nil)
}
