package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var Db *sql.DB

func init() {
	Db, err = sql.Open("mysql", "root:wangning@tcp(127.0.0.1:3306)/chitchat?parseTime=true")
	if err != nil {
		fmt.Println("sql connected fail")
	}
}