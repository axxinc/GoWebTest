package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Db      *sql.DB
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// post, err := retrieve(id)
	err = post.fetch(id)
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	len := r.ContentLength
	body := make([]byte, len)

	r.Body.Read(body)
	json.Unmarshal(body, &post)

	err = post.create()
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// post, err := retrieve(id)
	post.fetch(id)
	if err != nil {
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)

	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	// post, err := retrieve(id)
	post.fetch(id)
	if err != nil {
		return
	}

	err = post.delete()
	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}

// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	var err error

// 	switch r.Method {
// 	case "GET":
// 		err = handleGet(w, r)
// 	case "POST":
// 		err = handlePost(w, r)
// 	case "PUT":
// 		err = handlePut(w, r)
// 	case "DELETE":
// 		err = handleDelete(w, r)
// 	}
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

func handleRequest(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {

	var err error
	Db, err := sql.Open("postgres", "user=postgres password=123456 port=5432 host=127.0.0.1 sslmode=disable dbname=gwp")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/post/", handleRequest(&Post{Db: Db}))
	http.ListenAndServe(":8000", nil)
}
