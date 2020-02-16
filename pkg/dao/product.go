package dao

import (
	"github.com/jinzhu/gorm"
	DB "test/pkg/db"
	"test/src/constant"
)

func ProductListOrm() *gorm.DB {
	db := DB.ConnectDbOrm()
	db.AutoMigrate(&constant.Menu{})
	return db
}