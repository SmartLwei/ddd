package application

import (
	"ddd/infra/rdb"
)

type Factory struct {
	UserSvc *UserService
}

func NewController(repoFactory *rdb.Factory) *Factory {
	var c = Factory{}
	c.UserSvc = NewUserAppl(repoFactory.GetUserRepo())
	return &c
}

func (c *Factory) GetDemoController() *UserService {
	return c.UserSvc
}
