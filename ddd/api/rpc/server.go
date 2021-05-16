package rpc

import (
	"context"
	"net"

	gd "ddd/api/rpc/grpcdemo"
	"ddd/application"
	"ddd/conf"

	"google.golang.org/grpc"
)

type GRPCService struct {
	setting *conf.GrpcSetting
	server  *grpc.Server
	ctl     *application.Factory
}

func NewGrpcService(setting *conf.GrpcSetting, ctl *application.Factory) *GRPCService {
	var s GRPCService
	s.setting = setting
	s.ctl = ctl
	s.server = grpc.NewServer()
	return &s
}

func (gs *GRPCService) Run() {
	lis, err := net.Listen("tcp", gs.setting.Port)
	if err != nil {
		panic("failed to listen on port " + gs.setting.Port + ": " + err.Error())
	}
	gd.RegisterDemoServiceServer(gs.server, gs)
	if err := gs.server.Serve(lis); err != nil {
		panic("grpc server failed to serve:" + err.Error())
	}
}

func (gs *GRPCService) Stop() {
	gs.server.GracefulStop()
}

func (gs *GRPCService) GetDemos(ctx context.Context, req *gd.GetDemosReq) (*gd.GetDemosResp, error) {
	return gs.ctl.UserSvc.GrpcGetUsers(req)
}
