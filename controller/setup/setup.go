package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/lgyong511/sffld-go/config"
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
	cfg := new(config.Config)
	err := c.ShouldBindJSON(cfg)
	if err != nil {
		controller.Error(c, controller.ParamsErrorCode, err.Error())
		return
	}
	err = setup.UpdateSettings(cfg)
	if err != nil {
		controller.Error(c, controller.ErrorCode, err.Error())
		return
	}
	controller.Success(c, nil)
}
