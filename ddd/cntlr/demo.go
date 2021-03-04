package cntlr

import (
	"ddd/api/rest/dto"
	"ddd/infra/db"
)

type Demo struct{}

func NewDemo() *Demo {
	return &Demo{}
}

func (d *Demo) AddDemo(req *dto.AddDemoReq) error {
	domainDemo, err := req.ValidateAnd2Model()
	if err != nil {
		return err
	}
	return db.DemoRepo{}.Insert(db.GetTx(), domainDemo)
}

func (d *Demo) GetDemos(req *dto.GetDemosReq) (*dto.GetDemosResp, error) {
	filter, err := req.ValidateAnd2Model()
	if err != nil {
		return nil, err
	}
	count, domainDemos, err := db.DemoRepo{}.Get(db.GetTx(), filter)
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
