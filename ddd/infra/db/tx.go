package db

import (
	"ddd/conf"
	"ddd/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sync"
)

var tx Tx
var once = &sync.Once{}

type Tx interface {
	Begin() Tx
	Rollback() Tx
	Commit() Tx
	Error() error
	Gorm() *gorm.DB
}

type TxImp struct {
	db *gorm.DB
}

func Init() {
	once.Do(func() {
		db := newGorm()
		autoMigTables(db)
		tx = &TxImp{db: db}
	})
}

func newGorm() *gorm.DB {
	var dbSetting = conf.GetSetting().DB
	mysqlConfig := mysql.Config{
		DSN:                       dbSetting.URI,
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	gormConfig := gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,

		NamingStrategy: &schema.NamingStrategy{SingularTable: true},
		Logger:         logger.Default.LogMode(logger.LogLevel(dbSetting.LogLevel)),
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig)
	if err != nil {
		panic("db open failed: " + err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(dbSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dbSetting.MaxOpenConns)
	return db
}

func autoMigTables(db *gorm.DB) {
	if err := db.AutoMigrate(&domain.Demo{}); err != nil {
		panic("auto migrate demo failed: " + err.Error())
	}
}

func GetTx() Tx {
	if tx == nil {
		panic("get tx failed, please init tx first")
	}
	return tx
}

func (tx *TxImp) Gorm() *gorm.DB {
	return tx.db
}

func (tx *TxImp) Begin() Tx {
	return &TxImp{db: tx.Gorm().Begin()}
}

func (tx *TxImp) Rollback() Tx {
	return &TxImp{db: tx.Gorm().Rollback()}
}

func (tx *TxImp) Commit() Tx {
	return &TxImp{db: tx.Gorm().Commit()}
}

func (tx *TxImp) Error() error {
	return tx.db.Error
}
