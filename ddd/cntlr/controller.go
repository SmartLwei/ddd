package cntlr

import (
	"ddd/infra/db"
)

type Controller struct {
	DemoCtl *Demo
}

func NewController(repoFactory *db.Factory) *Controller {
	var c = Controller{}
	c.DemoCtl = NewDemo(repoFactory.GetDemoRepo())
	return &c
}

func (c *Controller) GetDemoController() *Demo {
	return c.DemoCtl
}
