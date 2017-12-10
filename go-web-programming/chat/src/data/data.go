

package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:wangning@tcp(127.0.0.1:3306)/chitchat")
	if err != nil {
		log.Fatal(err)
	}
	return
}