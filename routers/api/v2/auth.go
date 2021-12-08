package v2

import (
	"CLogServer/dao"
	"CLogServer/pkg/app"
	"CLogServer/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetToken 用户名、密码获取token
func GetToken(c *gin.Context) {
	var form dao.User
	appG := app.Gin{C: c}
	if err := c.ShouldBind(&form); err != nil {
		log.Println("参数绑定错误: ", err.Error())
		appG.Response(http.StatusBadRequest, app.Error, nil)
		return
	}

	user := dao.QueryByUsername(form.Username)
	if user == nil {
		log.Println("用户不存在")
		appG.Response(http.StatusOK, app.UserAuthError, nil)
		return
	}

	if user.Password != form.Password {
		log.Println("密码错误")
		appG.Response(http.StatusOK, app.UserAuthError, nil)
		return
	}

	token, err := util.GenerateToken(form.Username, form.Password)
	if err != nil {
		log.Println("内部服务错误")
		appG.Response(http.StatusInternalServerError, app.ServerError, nil)
		return
	}

	appG.Response(200, app.Success, token)
}
