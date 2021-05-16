package handler

import (
	"net/http"

	"ddd/api/rest/dto"
	"ddd/cntlr"

	"github.com/gin-gonic/gin"
)

type Demo struct {
	demoCtl *cntlr.Demo
}

func NewDemoHandler(demoCtl *cntlr.Demo) *Demo {
	return &Demo{demoCtl: demoCtl}
}

// AddDemo 数据库增接口
// @Summary 数据库增接口demo
// @description 数据库增接口demo的描述内容
// @tags demo
// @id InsertDemo
// @accept application/json
// @produce application/json
// @param data body dto.AddDemoReq false "AddDemoReq"
// @success 200 {object} dto.JSONResult{}
// @router /demo/api/v1/demo [post]
func (d *Demo) AddDemo(c *gin.Context) {

	// 获取用户请求
	var req dto.AddDemoReq
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, &dto.JSONResult{Code: 1, Message: "failed"})
		return
	}

	// 在 controller 层中处理请求
	err := d.demoCtl.AddDemo(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.JSONResult{Code: 2, Message: "failed"})
		return
	}

	// 向前端返回数据
	c.JSON(http.StatusOK, &dto.JSONResult{Message: "succeed"})
}

// GetDemos 数据库查接口
// @Summary 数据库查接口demo
// @description 数据库查接口demo的描述内容
// @tags demo
// @id GetDemos
// @accept application/json
// @produce application/json
// @param data query dto.GetDemosReq false "demoReq"
// @success 200 {object} dto.JSONResult{data=dto.GetDemosResp}
// @router /demo/api/v1/demos [get]
func (d *Demo) GetDemos(c *gin.Context) {

	// 获取请求参数
	var req dto.GetDemosReq
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, &dto.JSONResult{Code: 1, Message: "failed"})
		return
	}

	// 在 controller 层中处理请求
	resp, err := d.demoCtl.GetDemos(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.JSONResult{Code: 2, Message: "failed"})
		return
	}

	// 向前端返回数据
	c.JSON(http.StatusOK, &dto.JSONResult{Message: "succeed", Data: resp})
}
