package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Post struct {
	Id        int `gorm:"primary_key"`
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int    `sql:"index"`
	CreatedAt time.Time
}

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=123456 dbname=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Hello World", Author: "Sau Sheong"}
	fmt.Println(post)
	DB.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "Good Post", Author: "Joe"}
	DB.Model(&post).Association("Comments").Append(comment)

	var readPost Post
	DB.Where("author=$1", "Sau Sheong").First(&readPost)

	var comments []Comment
	DB.Model(&readPost).Related(&comments)

	fmt.Println(comments[0])
}
