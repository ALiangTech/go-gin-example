package models

type Auth struct {
	User_id  int    `gorm:"primary_key" json:"user_id"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("user_id").Where(Auth{Username: username, Password: password}).First(&auth)
	return auth.User_id > 0
}
