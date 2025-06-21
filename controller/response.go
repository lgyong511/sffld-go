package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	//成功
	SuccessCode = 666
	//失败
	ErrorCode = 700 + iota
	//参数错误
	ParamsErrorCode
	//未授权
	UnauthorizedCode
	//未登录
	UnauthenticatedCode
	//生成token失败
	GenTokenErrorCode
	//无token
	NoTokenCode
	//token无效
	InvalidTokenCode
	//用户名或密码错误
	UsernameOrPasswordErrorCode
)

// code转文字
func CodeToMsg(code int) string {
	switch code {
	case SuccessCode:
		return "success"
	case ErrorCode:
		return "error"
	case ParamsErrorCode:
		return "params error"
	case UnauthorizedCode:
		return "unauthorized"
	case UnauthenticatedCode:
		return "unauthenticated"
	case GenTokenErrorCode:
		return "gen token error"
	case UsernameOrPasswordErrorCode:
		return "username or password error"
	case NoTokenCode:
		return "no token"
	case InvalidTokenCode:
		return "invalid token"
	default:
		return "unknown"
	}
}

// api统一返回结构体，支持分页。
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// 定义成功返回
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, &Response{
		Code: SuccessCode,
		Msg:  "success",
		Data: data,
	})
}

// 定义失败返回
func Error(c *gin.Context, code int, msg string) {
	c.JSON(200, &Response{
		Code: code,
		Msg:  msg,
	})
}

// 定义未授权返回
func Unauthorized(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, &Response{
		Code: UnauthorizedCode,
		Msg:  msg,
	})
}

// 定义分页返回
func Page(c *gin.Context, data interface{}, count int) {
	c.JSON(200, &Response{
		Code: SuccessCode,
		Msg:  "success",
		Data: gin.H{
			"list":  data,
			"total": count,
		},
	})
}

// 定义自定义返回
func Custom(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// 定义token返回
func Token(c *gin.Context, token string) {
	c.JSON(200, &Response{
		Code: SuccessCode,
		Msg:  "success",
		Data: gin.H{
			"token": token,
		},
	})
}

// 定义退出登录返回
func Logout(c *gin.Context) {
	c.JSON(200, &Response{
		Code: SuccessCode,
		Msg:  "success",
	})
}

// 定义登录返回
func Login(c *gin.Context, token string) {
	c.JSON(200, &Response{
		Code: SuccessCode,
		Msg:  "success",
		Data: gin.H{
			"token": token,
		},
	})
}
