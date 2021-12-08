package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Error         = "error"
	Success       = "success"
	ServerError   = "内部服务错误"
	UserAuthError = "用户名或密码错误"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(code int, msg string, data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	return
}
