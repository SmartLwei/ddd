package cntlr

import "ddd/api/rest/dto"

type Demo struct{}

func NewDemo() *Demo {
	return &Demo{}
}

func (d *Demo) Demo(req *dto.DemoReq) (*dto.DemoResp, error) {
	return &dto.DemoResp{ID: req.ID}, nil
}
