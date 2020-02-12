package db

import (
	"github.com/jinzhu/gorm"
	"test/src/constant"
	"time"
)

//noinspection ALL
func ConnectDbOrm() *gorm.DB {
	db, err := gorm.Open("mysql", constant.DbConnection)
	db.DB().SetConnMaxLifetime(30 * time.Second)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(20)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}