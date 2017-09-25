package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:wangning@/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var (
		id int
		name string
	)

	rows, err := db.Query("select id, name from user where id = ?", 1)
	if err != nil {
		fmt.Println("err", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("err",err)
		}
		fmt.Println("result: ", id ,name)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("err", err)
	}
}