package routers

import (
	"CLogServer/middleware/jwt"
	v1 "CLogServer/routers/api/v1"
	v2 "CLogServer/routers/api/v2"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api1 := r.Group("/api/v1")
	api2 := r.Group("/api/v2")

	api1.Use(jwt.JWT())
	{
		api1.POST("/box", v1.AddBox)
	}

	api2.POST("/token", v2.GetToken)
	api2.Use(jwt.JWT())
	{

	}

	return r
}
