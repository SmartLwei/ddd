package rpc

import (
	"ddd/api/rpc/handler"
	"net"

	gd "ddd/api/rpc/dddcli"
	"ddd/application"
	"ddd/conf"

	"google.golang.org/grpc"
)

type GRPCService struct {
	setting *conf.GrpcSetting
	server  *grpc.Server
	handler *handler.GRPCHandler
}

func NewGrpcService(setting *conf.GrpcSetting, appFac *application.Factory) *GRPCService {
	var s GRPCService
	s.setting = setting
	s.server = grpc.NewServer()
	var hdl = handler.NewGRPCHandler(appFac)
	s.handler = hdl
	return &s
}

func (gs *GRPCService) Run() {
	lis, err := net.Listen("tcp", gs.setting.Port)
	if err != nil {
		panic("failed to listen on port " + gs.setting.Port + ": " + err.Error())
	}
	gd.RegisterDDDServiceServer(gs.server, gs.handler)
	if err := gs.server.Serve(lis); err != nil {
		panic("grpc server failed to serve:" + err.Error())
	}
}

func (gs *GRPCService) Stop() {
	gs.server.GracefulStop()
}
