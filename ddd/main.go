package main

import (
	"ddd/api/rest"
	"ddd/api/rest/handler"
	"ddd/cntlr"
	"ddd/conf"
	_ "ddd/docs"
	"ddd/infra/db"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var httpServer = rest.GetRouter()
	httpServer.Run()

	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGBUS)
	for sig := range chSig {
		fmt.Println("signal received:" + sig.String())
		httpServer.Stop()
		break
	}
	fmt.Println("service exited")
}

func init() {
	newInstances()
}

// 新建各层的实例，从底层一步步往上初始化各个实例
func newInstances() {
	conf.Init("ddd/conf/config.yaml")
	db.Init()
	cntlr.Init()
	handler.Init()
	rest.Init()
}
