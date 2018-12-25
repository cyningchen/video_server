package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddVideoDeletionRecord(vid string) (err error) {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec(video_id) VALUES(?)")
	if err != nil {
		return
	}
	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("Add video_del_rec failed: %v", err)
		return
	}
	defer stmtIns.Close()
	return
}
