package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// type Post struct {
// 	Id         int
// 	Content    string
// 	AuthorName string `db:`
// }
type Post struct {
	Id         int
	Content    string
	AuthorName string `db:"author"`
}

var DB *sqlx.DB

func init() {
	var err error
	DB, err = sqlx.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=gwp password=123456 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = DB.QueryRowx("select id, content, author from posts where id =$1", id).StructScan(&post)
	if err != nil {
		return
	}
	return
}
func (post *Post) Create() (err error) {
	err = DB.QueryRow("insert into posts (content, author) values ($1, $2)　returning id", post.Content, post.AuthorName).Scan(&post.Id)
	return
}
func main() {
	post := Post{Content: "Hello World!", AuthorName: "Sau Sheong"}
	post.Create()
	fmt.Println(post)

	post2, _ := GetPost(2)
	fmt.Println(post2)
}
