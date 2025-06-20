package main

import (
	"fmt"
	"time"

	"github.com/lgyong511/sffld-go/config"
	"github.com/lgyong511/sffld-go/config/lg"
	"github.com/lgyong511/sffld-go/config/vp"
	"github.com/sirupsen/logrus"
)

func main() {

	for {
		logrus.Debug("debug")
		logrus.Info("info")
		time.Sleep(10 * time.Second)
	}

}

func init() {
	vp := vp.New()

	conf := vp.Get()
	lg.SetLogurs(conf.Log)
	fmt.Printf("conf: %v\n", conf)

	vp.AddReloadCallback(func(conf *config.Config) {
		lg.SetLogurs(conf.Log)
		logrus.Info("配置文件变更，重新加载日志配置")
	})
}
