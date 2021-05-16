package main

import (
	"ddd/api/rest"
	"ddd/api/rpc"
	"ddd/application"
	"ddd/conf"
	_ "ddd/docs"
	"ddd/infra/rdb"
	"ddd/log"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	// global config
	log.Init("DEBUG")
	var setting = getSetting()
	log.InfoZ("get setting", zap.Reflect("setting", setting))

	// infrastructure layer
	var orm = rdb.NewGorm(setting.DB)
	var repoFactory = rdb.NewFactory(orm)

	// application layer
	var controller = application.NewController(repoFactory)

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
