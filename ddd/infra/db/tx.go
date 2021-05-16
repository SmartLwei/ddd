package db

import (
	"ddd/conf"
	"ddd/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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

// NewGorm new gorm connection pool
func NewGorm(setting *conf.DBSetting) *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       setting.URI,
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	gormConfig := gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,

		NamingStrategy: &schema.NamingStrategy{SingularTable: true},
		Logger:         logger.Default.LogMode(logger.LogLevel(setting.LogLevel)),
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig)
	if err != nil {
		panic("db open failed: " + err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(setting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(setting.MaxOpenConns)
	return db
}

func autoMigTables(db *gorm.DB) {
	if err := db.AutoMigrate(&domain.Demo{}); err != nil {
		panic("auto migrate demo failed: " + err.Error())
	}
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
