package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/lgyong511/sffld-go/controller/setup"
)

// RegisterSetupRouter 注册setup路由
func RegisterSetupRouter(rg *gin.RouterGroup) {
	settings(rg)
}

// settings 配置管理路由
func settings(rg *gin.RouterGroup) {
	rg.GET("/settings", setup.GetSettings)
	rg.POST("/settings", setup.UpdateSettings)
}
