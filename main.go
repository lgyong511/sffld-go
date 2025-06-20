package main

import (
	"time"

	"github.com/lgyong511/sffld-go/config"
	"github.com/lgyong511/sffld-go/config/lg"
	"github.com/lgyong511/sffld-go/config/vp"
	"github.com/sirupsen/logrus"
)

func main() {
	vp := vp.New()
	conf := vp.Get()
	logrus.WithFields(conf.ToLogFields()).Debug("配置信息")

	lg.SetLogurs(conf.Log)

	vp.AddReloadCallback(func(conf *config.Config) {
		lg.SetLogurs(conf.Log)
	})

	for {
		logrus.Debug("debug")
		logrus.Info("info")
		time.Sleep(10 * time.Second)
	}

}
