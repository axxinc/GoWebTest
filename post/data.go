package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=123456 port=5432 host=127.0.0.1 sslmode=disable dbname=gwp")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (post Post, err error) {
	post = Post{}
	err = DB.QueryRow("select id,content,author from posts where id=$1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) create() (err error) {
	err = DB.QueryRow("insert into posts (content,author) values ($1,$2) returning id", post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) update() (err error) {
	_, err = DB.Exec("update posts set content=$1,author=$2 where id=$3", post.Content, post.Author, post.Id)
	return
}

func (post *Post) delete() (err error) {
	_, err = DB.Exec("delete from posts where id=$1", post.Id)
	return
}
