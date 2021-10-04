package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Ctx context.Context
	Db  *sql.DB
)

func ConnectDb() {
	var err error

	var host string
	var pass string

	host = os.Getenv("MYSQL_HOST")
	pass = os.Getenv("MYSQL_ROOT_PASSWORD")

	fmt.Println("!!!!MYSQL host:" + host)

	if host == "" || pass == "" { // TODO remove add docker local
		host = "127.0.0.1:3306"
		pass = "test1234"
	}

	connStr := fmt.Sprintf("root:%s@(%s)/todo", pass, host)

	fmt.Println("!!!!", connStr)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	Db = db
}
