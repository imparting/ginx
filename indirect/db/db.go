package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	if db == nil {
		dsn := "root:phpts@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}
}

func HasTable(dst interface{}) bool {
	return db.Migrator().HasTable(dst)
}

func AutoMigrate(dst interface{}) error {
	return db.AutoMigrate(dst)
}

func GetDB() *gorm.DB {
	return db
}

func Table(tableName string) *gorm.DB {
	return db.Table(tableName)
}
