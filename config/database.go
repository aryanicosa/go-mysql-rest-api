package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(localhost:3306)/testdb")
}
