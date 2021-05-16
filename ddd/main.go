package main

import (
	"ddd/api/rest"
	"ddd/api/rpc"
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
	// global config
	var setting = getSetting()

	// infrastructure layer
	var orm = db.NewGorm(setting.DB)
	var repoFactory = db.NewFactory(orm)

	// application layer
	var controller = cntlr.NewController(repoFactory)

	// API layer
	var httpServer = rest.NewRouter(setting.Rest, controller)
	var grpcServer = rpc.NewGrpcService(setting.Grpc, controller)

	// run services
	httpServer.Run()
	grpcServer.Run()

	// wait signal to exit service
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGBUS)
	for sig := range chSig {
		fmt.Println("signal received:" + sig.String())
		httpServer.Stop()
		grpcServer.Stop()
		break
	}

	// gracefully exit service
	grpcServer.Stop()
	httpServer.Stop()

	fmt.Println("service exited")
}

func getSetting() *conf.Setting {
	return conf.BuildConfig("conf/config.yaml")
}
