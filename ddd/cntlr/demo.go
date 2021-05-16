package cntlr

import (
	"ddd/api/rest/dto"
	gd "ddd/api/rpc/grpcdemo"
	"ddd/domain"
	"ddd/infra/db"
)

type Demo struct {
	demoRepo db.DemoItf
}

// NewDemo new demo service to deal with demo model related service
func NewDemo(demoRepo db.DemoItf) *Demo {
	var demo = Demo{}
	demo.demoRepo = demoRepo
	return &demo
}

// AddDemo add demo to database
func (d *Demo) AddDemo(req *dto.AddDemoReq) error {
	domainDemo, err := req.ValidateAnd2Model()
	if err != nil {
		return err
	}
	return d.demoRepo.Insert(domainDemo)
}

// GetDemos get demos from database
func (d *Demo) GetDemos(req *dto.GetDemosReq) (*dto.GetDemosResp, error) {
	filter, err := req.ValidateAnd2Model()
	if err != nil {
		return nil, err
	}
	count, domainDemos, err := d.demoRepo.Get(filter)
	if err != nil {
		return nil, err
	}
	var dtoDemos []*dto.Demo
	for _, d := range domainDemos {
		var dd = &dto.Demo{}
		if err := dd.GenerateFromDomain(d); err != nil {
			return nil, err
		}
		dtoDemos = append(dtoDemos, dd)
	}
	return &dto.GetDemosResp{Count: count, Demos: dtoDemos}, nil
}

// GrpcGetDemos get demos from database for grpc interface
func (d *Demo) GrpcGetDemos(req *gd.GetDemosReq) (*gd.GetDemosResp, error) {
	var filter = &domain.DemoFilter{
		ID:     req.Id,
		Name:   req.Name,
		Offset: int(req.Offset),
		Limit:  int(req.Limit),
	}
	count, domainDemos, err := d.demoRepo.Get(filter)
	if err != nil {
		return nil, err
	}

	var grpcDemos []*gd.Demo
	for _, d := range domainDemos {
		var dd = &gd.Demo{
			Id:        d.ID,
			CreatedAt: d.CreatedAt,
			Name:      d.Name,
		}
		grpcDemos = append(grpcDemos, dd)
	}
	return &gd.GetDemosResp{Count: count, Demos: grpcDemos}, nil
}
