package modules

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Mysql() *sql.DB {
	dataSourceName := "zwhset:zuoloveyou@tcp(115.28.12.222:3306)/zwhset?charset=utf8"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	return db
}
