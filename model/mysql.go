package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xchao0213/go-demos/common"
)

var DB *sql.DB

const prefix = "home_"

func InitDB() (error){
	var dataSource = common.GetDataSource()
	common.Logger(dataSource)
	var err error

	DB, err = sql.Open("mysql",dataSource)
	common.CheckErr(err)

	DB.SetMaxIdleConns(20)
	DB.SetMaxOpenConns(20)

	return nil

}


