package db

import (
	"ddd/domain"
)

type DemoItf interface {
	Insert(tx Tx, demo *domain.Demo) error
	Get(tx Tx, filter *domain.DemoFilter) (int64, []*domain.Demo, error)
}

type DemoRepo struct{}

func (DemoRepo) Insert(tx Tx, demo *domain.Demo) error {
	return tx.Gorm().Create(demo).Error
}

func (DemoRepo) Get(tx Tx, filter *domain.DemoFilter) (int64, []*domain.Demo, error) {
	var db = tx.Gorm().Model(&domain.Demo{})
	var count int64
	var demos []*domain.Demo
	if filter.ID != 0 {
		db = db.Where("id=?", filter.ID)
	}
	if filter.Name != "" {
		db = db.Where("name=?", filter.Name)
	}
	db = db.Count(&count)
	if filter.Limit > 0{
		db = db.Limit(filter.Limit)
	}
	if filter.Offset > 0{
		db = db.Offset(filter.Offset)
	}
	db = db.Order("name desc")
	db = db.Find(&demos)
	return count, demos, db.Error
}
