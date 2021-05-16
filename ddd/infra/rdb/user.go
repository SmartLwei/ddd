package rdb

import (
	"ddd/domain"

	"gorm.io/gorm"
)

type UserItf interface {
	Insert(demo *domain.User) error
	Get(filter *domain.UserFilter) (int64, []*domain.User, error)
}

type UserRepo struct {
	orm *gorm.DB
}

func NewUserRepo(orm *gorm.DB) UserItf {
	return &UserRepo{orm: orm}
}

func (d *UserRepo) Insert(demo *domain.User) error {
	return d.orm.Create(demo).Error
}

func (d *UserRepo) Get(filter *domain.UserFilter) (int64, []*domain.User, error) {
	var db = d.orm.Model(&domain.User{})
	var count int64
	var demos []*domain.User
	if filter.ID != 0 {
		db = db.Where("id=?", filter.ID)
	}
	if filter.Name != "" {
		db = db.Where("name=?", filter.Name)
	}
	db = db.Count(&count)
	if filter.Limit > 0 {
		db = db.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		db = db.Offset(filter.Offset)
	}
	db = db.Order("name desc")
	db = db.Find(&demos)
	return count, demos, db.Error
}
