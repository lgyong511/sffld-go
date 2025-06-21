package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/lgyong511/sffld-go/controller/setup"
)

// RegisterSetupRouter 注册setup路由
func RegisterSetupRouter(rg *gin.RouterGroup) {
	setLog(rg)
	setApp(rg)
}

// setLog 设置日志
func setLog(rg *gin.RouterGroup) {
	rg.POST("/log", setup.SetLog)
}

// setApp 设置应用
func setApp(rg *gin.RouterGroup) {
	rg.POST("/app", setup.SetApp)
}
