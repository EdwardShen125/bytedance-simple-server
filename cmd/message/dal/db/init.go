package db

import (
	"github.com/EdwardShen125/bytedance-simple-server/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormtracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormtracing.New()); err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&MessageRaw{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(constants.MySQLMaxIdleConns)
	sqlDB.SetMaxOpenConns(constants.MySQLMaxOpenConns)
	sqlDB.SetConnMaxLifetime(constants.MySQLConnMaxLifetime)

}
