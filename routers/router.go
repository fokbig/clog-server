package routers

import (
	v1 "CLogServer/routers/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apis := r.Group("/api/v1")
	apis.Use()
	{
		apis.POST("/box", v1.AddBox)
	}

	return r
}
