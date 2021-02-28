package cntlr

import "sync"

type Controller struct {
	DemoCtl *Demo
}

var once = &sync.Once{}
var controller *Controller

func Init() {
	once.Do(func() {
		controller = &Controller{DemoCtl: &Demo{}}
	})
}

func GetController() *Controller {
	if controller == nil {
		panic("get controller failed, please init controller first")
	}
	return controller
}
