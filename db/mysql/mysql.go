package mysql

import (
	"database/sql"
	logger "github.com/sirupsen/logrus"
)

func NewDBConn(db *sql.DB){
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "goblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		logger.Error("Cann't connect to Mysql DB!: ", err.Error())
		panic(err.Error())
	}
	return
}