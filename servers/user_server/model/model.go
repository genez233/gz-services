package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gz-services/servers/user_server/global"
	"log"
	"os"
	"time"
)

func NewEngine() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond, // 慢 SQL 阈值
			LogLevel:                  logger.Info,      // Log level
			IgnoreRecordNotFoundError: true,             // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,             // 彩色打印
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		global.Database.UserName,
		global.Database.Password,
		global.Database.Host,
		global.Database.DBName,
		global.Database.Charset,
		global.Database.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	// 注册Callback
	err = db.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", disableQueryNotFound)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(global.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(global.Database.MaxIdleConns)

	return db, err
}

// 禁用查询NotFoundError
func disableQueryNotFound(db *gorm.DB) {
	db.Statement.RaiseErrorOnNotFound = false
}
