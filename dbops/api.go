package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
)


func AddUserCredential(loginName, pwd string) (err error) {
	// 预编译prepare
	stmtIns, err := dbConn.Prepare("INSERT INTO user(username, pwd) VALUES (?,?)")
	if err != nil{
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return
}

func GetUserCredential(loginName string) (pwd string, err error)  {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM user where username = ?")
	if err != nil{
		log.Println(err)
		return "", err
	}
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows{
		return "", err
	}
	defer stmtOut.Close()
	return
}

func DeleteUser(loginName string) (err error) {
	stmtDel, err := dbConn.Prepare("DELETE FROM user where username = ?")
	if err != nil{
		log.Println(err)
		return
	}
	_,  err = stmtDel.Exec(loginName)
	if err != nil{
		return
	}
	defer stmtDel.Close()
	return
}
