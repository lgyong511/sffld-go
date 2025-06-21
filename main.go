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

	r := gin.Default()
	r.Use(middleware.JwtAuth)

	router.RegisterRouter(r)
	g := graceful.New(":" + conf.App.Port)
	g.Start(r)

	for {
		logrus.Debug("debug")
		logrus.Info("info")
		time.Sleep(10 * time.Second)
	}

}
