package base

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var DB *gorm.DB

func init() {
	// https://github.com/go-sql-driver/mysql
	dbConf := ConfInstance.Db
	dsn := dbConf.Dsn
	if gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 驼峰表名，不加s
			SingularTable: true,
		},
	}); nil != err {
		panic(err)
	} else {
		if db, err := gdb.DB(); nil == err {
			db.SetMaxIdleConns(dbConf.MaxIdleConns)
			db.SetConnMaxLifetime(dbConf.ConnMaxLifetime * time.Second)
			db.SetConnMaxIdleTime(dbConf.ConnMaxIdleTime * time.Second)
			db.SetMaxOpenConns(dbConf.MaxOpenConns)
		} else {
			panic(err)
		}
		DB = gdb
	}
}
