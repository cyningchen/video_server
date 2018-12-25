package dbops

import "log"

func ReadVideoDeletionRecord(count int) (ids []string, err error) {
	stmtout, err := dbConn.Prepare("SELECT video_id FROM video_del_rec LIMIT ?")
	if err != nil {
		return
	}
	rows, err := stmtout.Query(count)
	if err != nil {
		log.Printf("Query video_del_rec error: %v", err)
		return
	}

	for rows.Next() {
		var id string
		if err := rows.Scan(id); err != nil {
			return
		}
		ids = append(ids, id)
	}
	defer stmtout.Close()
	return
}

func DelVideoDeletionRecord(vid string) (err error) {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_del_rec WHERE video_id=?")
	if err != nil {
		return
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("Deleteing video_del_rec error: %v", err)
		return
	}
	defer stmtDel.Close()
	return
}
