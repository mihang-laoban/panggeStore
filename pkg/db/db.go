package db

import (
	"database/sql"
	"log"
	"test/src/constant"
)

func ConnectDB() *sql.DB {
	db, err1 := sql.Open("mysql", constant.DbConnection)
	if err1 != nil{
		log.Fatalln(err1)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err2 := db.Ping(); err2 != nil{
		log.Fatalln(err2)
	}
	return db
}
