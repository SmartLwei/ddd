package dto

import "ddd/domain"

type JSONResult struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AddUserReq struct {
	Name string `json:"name" form:"name" example:"demo_name"`
}

func (req *AddUserReq) ValidateAnd2Model() (*domain.User, error) {
	return &domain.User{Name: req.Name}, nil
}

type User struct {
	ID        int64  `json:"id"`
	CreatedAt int64  `json:"created_at"`
	Name      string `json:"string"`
}

func (d *User) GenerateFromDomain(domainUser *domain.User) error {
	d.ID = domainUser.ID
	d.CreatedAt = domainUser.CreatedAt
	d.Name = domainUser.Name
	return nil
}

type GetUserReq struct {
	ID     int64  `json:"id" form:"id" example:"1"`
	Name   string `json:"name" form:"name" example:"demo_name"`
	Offset int    `json:"offset" form:"offset" example:"0"`
	Limit  int    `json:"limit" form:"limit" example:"10"`
}

func (req *GetUserReq) ValidateAnd2Model() (*domain.UserFilter, error) {
	return &domain.UserFilter{
		ID:     req.ID,
		Name:   req.Name,
		Offset: req.Offset,
		Limit:  req.Limit,
	}, nil
}

type GetUserResp struct {
	Count int64   `json:"count"`
	Demos []*User `json:"demos"`
}
