package mysql

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"log"
	"usercenter/config"
)

func GetDB() *sql.DB{
	config := conf.GetConfig().DB
	dsn := config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Name + "?charset=utf8&loc=Asia%2FShanghai"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	return  db
}

