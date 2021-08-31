package main

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// var DB *sql.DB
var DB *sqlx.DB

func init() {
	var err error
	// DB, err = sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=123456 dbname=gwp sslmode=disable")
	DB, err = sqlx.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=123456 dbname=gwp sslmode=disable")

	if err != nil {
		panic(err)
	}
}

type Post struct {
	Id      int
	Content string
	// Author   string
	AuthorName string `db:"author"`

	Comments []Comment
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

func (post *Post) Create() (err error) {
	err = DB.QueryRow("insert into posts(content,author) values ($1,$2) returning id", post.Content, post.AuthorName).Scan(&post.Id)
	return
}

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}
	err = DB.QueryRow("insert into comments(content,author,post_id) values ($1,$2,$3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	// err = DB.QueryRow("select id,content,author from posts where id=$1", id).Scan(&post.Id, &post.Content, &post.Author)
	err = DB.QueryRowx("select id,content,author from posts where id=$1", id).StructScan(&post)

	if err != nil {
		return
	}

	rows, err := DB.Query("select id,content,author from comments where post_id=$1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)

	}
	rows.Close()
	return
}

func main() {
	// post := Post{Content: "Hello World", Author: "Sau Sheong"}
	// post.Create()
	// fmt.Println(post)
	// comment := Comment{Content: "Good Post", Author: "Joe", Post: &post}
	// comment.Create()
	// fmt.Println(comment)
	post, _ := GetPost(1)
	fmt.Println(post)

}
