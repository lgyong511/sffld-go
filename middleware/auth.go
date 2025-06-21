package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lgyong511/sffld-go/controller"
	"github.com/lgyong511/sffld-go/util/jwt"
)

const (
	// 登录路由
	LoginPath = "/api/login"
)

// JwtAuth 鉴权中间件
func JwtAuth(c *gin.Context) {
	// 排除登录路由
	if strings.HasPrefix(c.Request.URL.Path, LoginPath) {
		c.Next()
		return
	}
	// 获取token
	token := c.Request.Header.Get("Authorization")
	// 校验token
	if token == "" {
		controller.Unauthorized(c, controller.CodeToMsg(controller.NoTokenCode))
		return
	}
	// 分离Bearer
	token = strings.TrimPrefix(token, "Bearer ")
	// 解析token
	claims, err := jwt.ParseToken(token)
	if err != nil {
		controller.Unauthorized(c, controller.CodeToMsg(controller.InvalidTokenCode))
		return
	}
	// 保存token信息到上下文
	c.Set("claims", claims)
	c.Next()
}
