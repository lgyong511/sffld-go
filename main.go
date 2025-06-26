package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lgyong511/graceful"
	"github.com/lgyong511/sffld-go/config"
	"github.com/lgyong511/sffld-go/config/lg"
	"github.com/lgyong511/sffld-go/config/vp"
	"github.com/lgyong511/sffld-go/middleware"
	"github.com/lgyong511/sffld-go/router"
	"github.com/lgyong511/sffld-go/util/jwt"
	"github.com/sirupsen/logrus"
)

func main() {
	mv := vp.New()
	vp.SetViperCallback(func() *vp.MgrViper {
		return mv
	})
	conf := mv.Get()
	lg.SetLogurs(conf.Log)
	jwt.SetTimeout(conf.App.AuthTimeout)

	logrus.WithFields(conf.ToConfigFields()).Debug("配置信息")

	//配置热更新后，重新配置logrus
	mv.AddReloadCallback(func(conf *config.Config) {
		lg.SetLogurs(conf.Log)
	})

	r := gin.Default()
	r.Use(middleware.JwtAuth)

	router.RegisterRouter(r)
	g := graceful.New(":" + conf.App.Port)
	g.Start(r)

	// 配置热更新后，重启gin
	mv.AddReloadCallback(func(conf *config.Config) {
		g.Restart(":"+conf.App.Port, nil)
	})

	for {
		logrus.Debug("debug")
		logrus.Info("info")
		time.Sleep(20 * time.Second)
	}

}
