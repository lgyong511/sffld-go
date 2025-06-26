package api

import "github.com/gin-gonic/gin"

// RegisterAPIRouter 注册api路由
func RegisterAPIRouter(rg *gin.RouterGroup) {
	login(rg)
	register(rg)
	logout(rg)
}
