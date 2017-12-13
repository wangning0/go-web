package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:wangning@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		panic(err)
	}
}


func retrive(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = ?", id).
		Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) create() (err error) {
	rs, err := Db.Exec("insert into posts(content, author) values (?, ?)", post.Content, post.Author)
	if err != nil {
		return 
	}
	id, err := rs.LastInsertId()
	post.Id = int(id)
	return
}

func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set content = ?, author = ? where id = ?", post.Content, post.Author, post.Id)
	return
}

func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = ?", post.Id)
	return
}