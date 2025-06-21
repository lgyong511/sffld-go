package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lgyong511/sffld-go/router/api"
	"github.com/lgyong511/sffld-go/router/setup"
)

// 路由

// RegisterGinRouter 注册gin路由
func RegisterRouter(r *gin.Engine) {
	apiGroup := r.Group("/api")
	api.RegisterAPIRouter(apiGroup)

	setupGroup := r.Group("/setup")
	setup.RegisterSetupRouter(setupGroup)
}
