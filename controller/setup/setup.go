package setup

import "github.com/gin-gonic/gin"

// 日志配置handler
func SetLog(c *gin.Context) {
	c.JSON(200, gin.H{
		"log": "debug",
	})
}

// 应用配置handler
func SetApp(c *gin.Context) {
	c.JSON(200, gin.H{
		"app": "debug",
	})
}
