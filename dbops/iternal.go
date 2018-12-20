package dbops

import "strconv"

func InsertSession(sid string, ttl int64, uname string) (err error) {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions(session_id, TTL, login_name) VALUES (?,?,?)")
	if err != nil {
		return
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return
	}
	return
}
