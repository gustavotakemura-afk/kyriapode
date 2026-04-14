package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	return sql.Open("mysql", "root:escola@tcp(localhost:3306)/mobile_trab1")
}
