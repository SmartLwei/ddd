package handler

import (
	"context"
	gd "ddd/api/rpc/dddcli"
	"ddd/application"
)

type GRPCHandler struct {
	ss *application.Factory
}

func NewGRPCHandler(ss *application.Factory) *GRPCHandler {
	return &GRPCHandler{ss: ss}
}

func (gs *GRPCHandler) GetUsers(ctx context.Context, req *gd.GetUsersReq) (*gd.GetUsersResp, error) {
	return gs.ss.UserSvc.GrpcGetUsers(req)
}
