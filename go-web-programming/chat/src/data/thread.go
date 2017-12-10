package data

import (
	"time"
)

type Thread struct {
	Id int
	Uuid string
	Topic string
	UserId int
	CreateAt time.Time
}

type Post struct {
	Id int
	Uuid string
	Body string
	UserId int
	ThreadId int
	CreateAt time.Time
}

// NumReplies
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("select count(*) from posts where thread_id = ?", thread.Id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	return
}

// CreatedAtDate
func (thread *Thread) CreatedAtDate() string {
	return thread.CreateAt.Format("Jan 2, 2006 at 3:04pm")
}

// User
func (thread *Thread) User() (user User) {
	user = User{}
	Db.QueryRow("select id, uuid, name, email, created_at from users where id = ?", thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreateAt)
	return
}

// posts
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("select id, uuid, body, user_id, thread_id, created_at from posts where thread_id = ?", thread.Id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreateAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	return
}

func (post *Post) CreatedAtDate() string {
	return post.CreateAt.Format("Jan 2, 2006 at 3:04pm")
}

func (post *Post) User() (user User) {
	user = User{}
	Db.QueryRow("select id, uuid, name, email, created_at from users where id = ?", post.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreateAt)
	return
}

// 获取所有的threads
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("select id, uuid, topic, user_id, created_at from threads")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreateAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	return
}

func ThreadByUUID(uuid string) (conv Thread, err error){
	conv = Thread{}
	err = Db.QueryRow("select id, uuid, topic, user_id, created_at from threads where uuid = ?", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreateAt)
	return
}