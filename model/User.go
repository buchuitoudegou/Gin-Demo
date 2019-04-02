package model

func UserAuth(username string, pwd string) bool {
	if username == "admin" && pwd == "123456" {
		return true
	} else {
		return false
	}
}

func UserRegister(username string, pwd string) string {
	if username == "admin" {
		return "multiple username"
	} else if len(pwd) < 6 {
		return "Password is too short"
	} else {
		return "OK"
	}
}