package utils

//数据库连接
import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//全局变量Db
var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:010911@tcp(localhost:3306)/bookstore")
	if err != nil {
		panic(err.Error())
	}
}

func main() {

}
