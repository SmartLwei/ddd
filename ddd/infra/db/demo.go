package db

import (
	"ddd/domain"

	"gorm.io/gorm"
)

type DemoItf interface {
	Insert(demo *domain.Demo) error
	Get(filter *domain.DemoFilter) (int64, []*domain.Demo, error)
}

type DemoRepo struct {
	orm *gorm.DB
}

func NewDemo(orm *gorm.DB) DemoItf {
	return &DemoRepo{orm: orm}
}

func (d *DemoRepo) Insert(demo *domain.Demo) error {
	return d.orm.Create(demo).Error
}

func (d *DemoRepo) Get(filter *domain.DemoFilter) (int64, []*domain.Demo, error) {
	var db = d.orm.Model(&domain.Demo{})
	var count int64
	var demos []*domain.Demo
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
