package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	// var dataSource = "root:123456@tcp(localhost:3306)/home?charset=latin1"
	// DB,err = sql.Open("mysql",dataSource)

	// if err != nil {

	// }
}


