package v1

import (
	"CLogServer/pkg/app"
	"CLogServer/service/box_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Box struct {
	From     string   `json:"from"`
	Filename string   `json:"filename"`
	Lines    []string `json:"lines"`
}

// AddBox 新增Box数据
func AddBox(c *gin.Context) {
	appG := app.Gin{C: c}
	var box Box
	if err := c.ShouldBindJSON(&box); err != nil {
		appG.Response(http.StatusBadRequest, err.Error(), nil)
		return
	}
	boxService := box_service.Box{
		From:     box.From,
		Filename: box.Filename,
		Lines:    box.Lines,
	}
	boxService.ReceiveBox()
	appG.Response(http.StatusOK, "", nil)
}
