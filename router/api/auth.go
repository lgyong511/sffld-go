package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lgyong511/sffld-go/controller/api"
)

// signIn 登录
func login(rg *gin.RouterGroup) {
	rg.POST("/login", api.Login)

}

// signUp 注册
func register(rg *gin.RouterGroup) {
	rg.POST("/register", api.Register)
}

// signOut 注销
func logout(rg *gin.RouterGroup) {
	rg.POST("/logout", api.Logout)
}
