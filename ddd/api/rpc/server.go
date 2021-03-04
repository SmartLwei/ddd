package rpc

import (
	"context"
	gd "ddd/api/rpc/grpcdemo"
	"ddd/conf"
	"ddd/domain"
	"ddd/infra/db"
	"google.golang.org/grpc"
	"net"
	"sync"
)

type GRPCService struct {
	setting *conf.GrpcSetting
	server  *grpc.Server
}

var once sync.Once
var grpcService *GRPCService

func Init() {
	once.Do(func() {
		grpcSetting := conf.GetSetting().Grpc
		grpcService = &GRPCService{setting: grpcSetting}
	})
}

func GetGRPCService() *GRPCService {
	if grpcService == nil {
		panic("get grpc service failed, please init first")
	}
	return grpcService
}

func (gs *GRPCService) Run() {
	lis, err := net.Listen("tcp", gs.setting.Port)
	if err != nil {
		panic("failed to listen on port " + gs.setting.Port + ": " + err.Error())
	}
	gs.server = grpc.NewServer()
	gd.RegisterDemoServiceServer(gs.server, gs)
	if err := gs.server.Serve(lis); err != nil {
		panic("grpc server failed to serve:" + err.Error())
	}
}

func (gs *GRPCService) Stop() {
	gs.server.GracefulStop()
}

func (gs *GRPCService) GetDemos(ctx context.Context, req *gd.GetDemosReq) (*gd.GetDemosResp, error) {
	filter := &domain.DemoFilter{ID: req.Id, Name: req.Name, Offset: int(req.Offset), Limit: int(req.Limit)}
	count, domainDemos, err := db.DemoRepo{}.Get(db.GetTx(), filter)
	var dtoDemos []*gd.Demo
	for _, d := range domainDemos {
		var demo = &gd.Demo{
			Id:        int64(d.ID),
			CreatedAt: d.CreatedAt,
			Name:      d.Name,
		}
		dtoDemos = append(dtoDemos, demo)
	}
	return &gd.GetDemosResp{Count: count, Demos: dtoDemos}, err

}
