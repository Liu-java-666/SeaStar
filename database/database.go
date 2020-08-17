package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Open(auth, pwd, addr, dbname string, port int) error {
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		auth, pwd, addr, port, dbname)
	db_, err := sqlx.Open("mysql", connstr)
	if err != nil {
		return err
	}
	db = db_

	return nil
}

func Select(dest interface{}, query string, args ...interface{}) error {
	return db.Select(dest, query, args...)
}

func Get(dest interface{}, query string, args ...interface{}) error {
	return db.Get(dest, query, args...)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}