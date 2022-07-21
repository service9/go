package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//全局变量
var (
	Db  *sql.DB
	err error
)

func init() { //赋值Db
	Db, err = sql.Open("mysql", "root:010911@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}

}
