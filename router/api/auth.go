package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lgyong511/sffld-go/controller/api"
)

// login 登录
func login(rg *gin.RouterGroup) {
	rg.POST("/login", api.Login)

}
