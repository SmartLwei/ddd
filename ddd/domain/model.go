package domain

import (
	"gorm.io/gorm"
)

type Model struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
