package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/lgyong511/sffld-go/controller"
	"github.com/lgyong511/sffld-go/service/setup"
)

// GetSettings 获取配置
func GetSettings(c *gin.Context) {
	cfg := setup.GetSettings()
	controller.Success(c, cfg)
}

// UpdateSettings 设置配置
func UpdateSettings(c *gin.Context) {
	c.JSON(200, gin.H{
		"setup": "debug",
	})
}
