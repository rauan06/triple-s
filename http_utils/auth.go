package http_utils

import . "triples/bucket_struct"

func Logout() {
	SessionUser = nil
}

func Login(token string) {
	if SessionUser != nil && token == "" {
		return
	}

	var tempUser *User
	for _, user := range AllUsers {
		if user.UserID == token {
			tempUser = user
			break
		}
	}

	if tempUser == nil && token != "" {
		tempUser = &User{UserID: token, Username: "cookie"}
	}

	AllUsers = append(AllUsers, tempUser)

	SessionUser = tempUser
}
