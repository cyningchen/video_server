package dbops

import "testing"

func clearTables()  {
	dbConn.Exec("truncate user")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
}

func TestUserWorkFlow(t *testing.T)  {
	t.Run("Add", TestAddUserCredential)
	t.Run("Get", TestGetUserCredential)
        t.Run("Del", TestDeleteUser)

}

func TestAddUserCredential(t *testing.T) {
	err := AddUserCredential("chenxl", "123")
	if err != nil{
		t.Errorf("Error of AddUser: %v", err)
	}
}

func TestGetUserCredential(t *testing.T) {
	_, err := GetUserCredential("chenxl")
	if err != nil{
		t.Errorf("Error of GetUser: %v", err)
	}

}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("chenxl")
	if err != nil{
		t.Errorf("Error of Del user: %v", err)
	}
}
