package mysql

import (
	"log"
)

type User struct {
	Uid int
	Username string
	Password string
}

func AddUser(username string, password string) error {
	_, err := GetDB().Exec("insert into user (username, password) values (?,?)", username, password)
	return err
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := GetDB().QueryRow("select * from user where username = ?", username).Scan(&user.Uid, &user.Username, &user.Password)
	if err != nil {
		log.Fatalln(err)
	}
	return &user, err
}

func GetUserByUid(uid int) (*User, error) {
	var user User
	err := GetDB().QueryRow("select * from user where uid = ?", uid).Scan(&user.Uid, &user.Username, &user.Password)
	if err != nil {
		log.Fatalln(err)
	}
	return  &user, err
}

