package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)


var (
	dbConn *sql.DB
	err error
)

func init()  {
	dbConn , err = sql.Open("mysql", "kedacom:KedaMysql16#@tcp(139.196.90.101:3306)/video_server?charset=utf8")
	if err != nil{
		panic(err.Error())
	}
}