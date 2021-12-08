package jwt

import (
	"CLogServer/pkg/app"
	"CLogServer/pkg/util"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		appG := app.Gin{C: c}
		code = 200
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = 401
		} else {
			token = token[7:]
			_, err := util.ParseToken(token)
			if err != nil {
				code = 401
			}
		}

		if code != 200 {
			appG.Response(10001, "签名错误", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
