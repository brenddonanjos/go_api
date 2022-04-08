package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	DB_NAME = "go_api_db"
	DB_USER = "user"
	DB_PASS = "user"
	DB_HOST = "mysql_app"
)

func Open() *sqlx.DB {
	db, err := sqlx.Connect("mysql", DB_USER+":"+DB_PASS+"@(mysql:3306)/"+DB_NAME+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	return db
}
