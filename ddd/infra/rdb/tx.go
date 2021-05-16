package rdb

import (
	"gorm.io/gorm"
)

type Tx interface {
	Begin() Tx
	Rollback() Tx
	Commit() Tx
	Error() error
	Orm() *gorm.DB
}

type TxImp struct {
	db *gorm.DB
}

func (tx *TxImp) Orm() *gorm.DB {
	return tx.db
}

func (tx *TxImp) Begin() Tx {
	return &TxImp{db: tx.Orm().Begin()}
}

func (tx *TxImp) Rollback() Tx {
	return &TxImp{db: tx.Orm().Rollback()}
}

func (tx *TxImp) Commit() Tx {
	return &TxImp{db: tx.Orm().Commit()}
}

func (tx *TxImp) Error() error {
	return tx.db.Error
}
