package controller

func LoginHandler(user string, pwd string) bool {
	if user == "admin" && pwd == "123456" {
		return true
	} else {
		return false
	}
}

func RegisterHandler(user string, pwd string) string {
	if user == "admin" {
		return "multiple username"
	} else if len(pwd) < 6 {
		return "length of pwd is too short"
	} else {
		return "OK"
	}
}