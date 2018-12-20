package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"video_server/defs"
	"video_server/utils"
)

func AddUserCredential(loginName, pwd string) (err error) {
	// 预编译prepare
	stmtIns, err := dbConn.Prepare("INSERT INTO user(username, pwd) VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return
}

func GetUserCredential(loginName string) (pwd string, err error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM user where username = ?")
	if err != nil {
		log.Println(err)
		return "", err
	}
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return
}

func DeleteUser(loginName string) (err error) {
	stmtDel, err := dbConn.Prepare("DELETE FROM user where username = ?")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = stmtDel.Exec(loginName)
	if err != nil {
		return
	}
	defer stmtDel.Close()
	return
}

// video

func AddNewVideo(aid int, name string) (video *defs.VideoInfo, err error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return
	}
	ctime := time.Now().Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare("INSERT INTO video_info(id, author_id, name, display_ctime) VALUES (?,?,?,?,?)")
	if err != nil {
		return
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return
	}
	defer stmtIns.Close()
	video = &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	return
}

func GetVideoInfo(vid string) (video *defs.VideoInfo, err error) {
	stmtOut, err := dbConn.Prepare("SELECT * FROM video_info where id = ?")
	if err != nil {
		log.Println(err)
		return
	}
	var aid int
	var dct string
	var name string
	err = stmtOut.QueryRow(vid).Scan(&aid, &dct, &name)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer stmtOut.Close()
	video = &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}
	return
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}
