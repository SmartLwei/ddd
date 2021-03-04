package dto

import "ddd/domain"

type JSONResult struct {
	Code    int64 `json:"code"`
	Data    interface{}
	Message string `json:"message"`
}

type AddDemoReq struct {
	Name string `json:"name" form:"name" example:"demo_name"`
}

func (req *AddDemoReq) ValidateAnd2Model() (*domain.Demo, error) {
	return &domain.Demo{Name: req.Name}, nil
}

type Demo struct {
	ID        uint   `json:"id"`
	CreatedAt int64  `json:"created_at"`
	Name      string `json:"string"`
}

func (d *Demo) GenerateFromDomain(domainDemo *domain.Demo) error {
	d.ID = domainDemo.ID
	d.CreatedAt = domainDemo.CreatedAt
	d.Name = domainDemo.Name
	return nil
}

type GetDemosReq struct {
	ID     int64  `json:"id" form:"id" example:"1"`
	Name   string `json:"name" form:"name" example:"demo_name"`
	Offset int    `json:"offset" form:"offset" example:"0"`
	Limit  int    `json:"limit" form:"limit" example:"10"`
}

func (req *GetDemosReq) ValidateAnd2Model() (*domain.DemoFilter, error) {
	return &domain.DemoFilter{
		ID:     req.ID,
		Name:   req.Name,
		Offset: req.Offset,
		Limit:  req.Limit,
	}, nil
}

type GetDemosResp struct {
	Count int64   `json:"count"`
	Demos []*Demo `json:"demos"`
}
