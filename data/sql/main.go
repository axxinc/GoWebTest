package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=gwp password=123456 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (post *Post) Create() (err error) {
	statment := "insert into posts (content,author) values ($1,$2) returning id"
	stmt, err := DB.Prepare(statment)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	if err != nil {
		return
	}
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	DB.QueryRow("select id,content,author from posts where id=$1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Update() (err error) {
	_, err = DB.Exec("update posts set content=$2,author=$3 where id=$1", post.Id, post.Content, post.Author)
	return
}

func Delete(id int) (err error) {
	_, err = DB.Exec("delete from posts where id=$1", id)
	return
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := DB.Query("select id,content,author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "zhangsan"}
	post.Create()
	fmt.Println(post)

	readPost, err := GetPost(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	readPost, err = GetPost(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(readPost)

	// Delete(3)
	posts, _ := Posts(5)
	for _, post := range posts {
		fmt.Println(post)
	}
}
