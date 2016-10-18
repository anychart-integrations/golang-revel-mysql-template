package controllers

import (
	"github.com/go-gorp/gorp"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Dbm *gorp.DbMap
)

func InitDB(){
	db, err := sql.Open("mysql", "anychart_user:anychart_pass@/anychart_db")
	if err != nil {
		panic(err)
	}
	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
}