package application

import (
	"ddd/api/rest/dto"
	gd "ddd/api/rpc/grpcdemo"
	"ddd/domain"
	"ddd/infra/rdb"
)

// UserService deal user related request or event
type UserService struct {
	userRepo rdb.UserItf
}

// NewUserAppl new user service to deal with user model related service
func NewUserAppl(demoRepo rdb.UserItf) *UserService {
	var userCtl = UserService{}
	userCtl.userRepo = demoRepo
	return &userCtl
}

// AddUser add user to system
func (d *UserService) AddUser(req *dto.AddUserReq) error {
	domainDemo, err := req.ValidateAnd2Model()
	if err != nil {
		return err
	}
	return d.userRepo.Insert(domainDemo)
}

// GetUsers get users from system
func (d *UserService) GetUsers(req *dto.GetUserReq) (*dto.GetUserResp, error) {
	filter, err := req.ValidateAnd2Model()
	if err != nil {
		return nil, err
	}
	count, domainUsers, err := d.userRepo.Get(filter)
	if err != nil {
		return nil, err
	}
	var dtoUsers []*dto.User
	for _, d := range domainUsers {
		var dd = &dto.User{}
		if err := dd.GenerateFromDomain(d); err != nil {
			return nil, err
		}
		dtoUsers = append(dtoUsers, dd)
	}
	return &dto.GetUserResp{Count: count, Demos: dtoUsers}, nil
}

// GrpcGetUsers get demos from database for grpc interface
func (d *UserService) GrpcGetUsers(req *gd.GetDemosReq) (*gd.GetDemosResp, error) {
	var filter = &domain.UserFilter{
		ID:     req.Id,
		Name:   req.Name,
		Offset: int(req.Offset),
		Limit:  int(req.Limit),
	}
	count, domainDemos, err := d.userRepo.Get(filter)
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
