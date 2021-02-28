package handler

import (
	"ddd/api/rest/dto"
	"ddd/cntlr"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Demo struct {
	demoCtl *cntlr.Demo
}

func NewDemoHandler(demoCtl *cntlr.Demo) *Demo {
	return &Demo{demoCtl: demoCtl}
}

// Demo demo
// @Summary demo接口功能
// @description demo接口功能描述
// @tags demo
// @id Demo
// @accept application/json
// @produce application/json
// @param data query dto.DemoReq false "demoReq"
// @success 200 {object} dto.JSONResult{data=dto.DemoResp}
// @router /svc-name/api/v1/demo/hello [get]
func (d *Demo) Demo(c *gin.Context) {

	// 获取请求参数
	var req dto.DemoReq
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// 在 controller 层中处理请求
	resp, err := d.demoCtl.Demo(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// 向前端返回数据
	c.JSON(http.StatusOK, resp)
}
