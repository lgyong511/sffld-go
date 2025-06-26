package setup

import (
	"github.com/lgyong511/sffld-go/config"
	"github.com/lgyong511/sffld-go/config/vp"
)

// GetSetup 获取配置
func GetSettings() *config.Config {
	mv := vp.Callback()
	return mv.Get()
}

// UpdateSettings 更新配置
func UpdateSettings(cfg *config.Config) error {
	mv := vp.Callback()
	mv.Set(cfg)
	return mv.Write()
}
