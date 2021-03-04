package handler

import (
	"ddd/cntlr"
	"sync"
)

type Handler struct {
	Demo *Demo
}

var once = &sync.Once{}
var handler *Handler

func Init() {
	once.Do(func() {
		handler = &Handler{Demo: &Demo{demoCtl: cntlr.GetController().DemoCtl}}
	})
}

func GetHandler() *Handler {
	if handler == nil {
		panic("get handler failed, please init handler first")
	}
	return handler
}
