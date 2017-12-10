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
	"time"
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