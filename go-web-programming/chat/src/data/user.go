// package data

// import (
// 	"time"
// )

// type User struct {
// 	Id int
// 	Uuid string
// 	Name string
// 	Email string
// 	Password string
// 	CreateAt time.Time
// }

// type Session struct {
// 	Id int
// 	Uuid string
// 	Email string
// 	UserId string
// 	CreateAt time.Time
// }

// func (session *Session) Check() (valid bool, err error) {
// 	err = Db.QueryRow("select id, uuid, email, user_id, created_at from sessions where uuid = ?", session.Uuid).
// 		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreateAt)
// 	if err != nil {
// 		valid = false
// 		return
// 	}
// 	if session.Id != 0 {
// 		valid = true
// 	}
// 	return
// }

package data

import (
	"fmt"
	"time"
	"errors"
	"database/sql"
)

type User struct {
	Id int
	Uuid string
	Name string
	Email string
	Password string
	CreateAt time.Time
}

type Session struct {
	Id int
	Uuid string
	Email string
	UserId int
	CreateAt time.Time
}

func (sess *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("select id, uuid, email, user_id, created_at from sessions where uuid = ?", sess.Uuid).
		Scan(&sess.Id, &sess.Uuid, &sess.Email, &sess.UserId, &sess.CreateAt)
	if err != nil {
		valid = false
		return
	}
	if sess.Id != 0 {
		valid = true
	}
	return
}

func (sess *Session) DeleteByUUID() (err error) {
	_, err = Db.Exec("delete from sessions where uuid = ?", sess.Uuid)
	if err != nil {
		return
	}
	return
}

func (sess *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, uuid, name, email, created_at from users where id = ?", sess.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreateAt)
	return
}

func (user *User) Create() (err error) {
	var name string
	fmt.Println(user.Name)
	err = Db.QueryRow("select id from users where name = ?", user.Name).
		Scan(&name)
	fmt.Println(name)
	if name != "" {
		return errors.New("user exist")
	}
	_, err = Db.Exec("insert into users (uuid, name, email, password, created_at) values (?, ?, ?, ?, ?)", createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now())
	if err != nil {
		return
	}
	return
}

func (user *User) CreateSession() (sess Session, err error) {
	var uuid string
	var rs sql.Result
	var t time.Time
	var id int64
	uuid = createUUID()
	t = time.Now()
	rs, err = Db.Exec("insert into sessions (uuid, email, user_id, created_at) values (?, ?, ?, ?)", uuid, user.Email, user.Id, t)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	sess = Session{
		Id: int(id),
		Uuid: uuid,
		Email: user.Email,
		UserId: user.Id,
		CreateAt: t,
	}
	return
}

func (user *User) CreateThread(topic string) (conv Thread, err error) {
	var uuid string
	var rs sql.Result
	var t time.Time
	var id int64
	uuid = createUUID()
	t = time.Now()
	rs, err = Db.Exec("insert into threads (uuid, topic, user_id, created_at) values (?, ?, ?, ?)", uuid, topic, user.Id, t)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	conv = Thread{
		Id: int(id),
		Uuid: uuid,
		Topic: topic,
		UserId: user.Id,
		CreateAt: t,
	}
	return
}

func (user *User) CreatePost(thread Thread, body string) (post Post, err error) {
	var uuid string
	var rs sql.Result
	var t time.Time
	var id int64
	uuid = createUUID()
	t = time.Now()
	rs, err = Db.Exec("insert into posts (uuid, body, user_id, thread_id, created_at) values (?, ?, ?, ?, ?)", uuid, body, user.Id, thread.Id, t)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	post = Post{
		Id: int(id),
		Uuid: uuid,
		Body: body,
		UserId: user.Id,
		ThreadId: thread.Id,
		CreateAt: t,
	}
	return
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, name, email, password, created_at from users where email = ?", email).
		Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreateAt)
	return
}