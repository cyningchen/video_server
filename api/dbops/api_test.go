package dbops

import "testing"

var tempvid string

func clearTables() {
	dbConn.Exec("truncate user")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", TestAddUserCredential)
	t.Run("Get", TestGetUserCredential)
	t.Run("Del", TestDeleteUser)

}

func TestAddUserCredential(t *testing.T) {
	err := AddUserCredential("chenxl", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func TestGetUserCredential(t *testing.T) {
	_, err := GetUserCredential("chenxl")
	if err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}

}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("chenxl")
	if err != nil {
		t.Errorf("Error of Del user: %v", err)
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", TestAddUserCredential)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}
