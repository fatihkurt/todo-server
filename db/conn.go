package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Ctx context.Context
	Db  *sql.DB
)

func ConnectDb() {
	var err error
	db, err := sql.Open("mysql", "root:password@/todo")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	Db = db
}
