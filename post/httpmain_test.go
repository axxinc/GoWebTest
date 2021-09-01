package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer = httptest.NewRecorder()
}

func TestHandleGet1(t *testing.T) {
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Error("Response code is ", writer.Code)
	}

	var post Post
	data := writer.Body.Bytes()
	json.Unmarshal(data, &post)
	if post.Id != 1 {
		t.Error("Post id is", post.Id)
	}
}

func TestHandlePut1(t *testing.T) {
	json := strings.NewReader(`{"content":"Update Post by 91","author":"91"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)

	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Error("Response code is ", writer.Code)
	}

	post, _ := retrieve(1)
	if post.Author != "91" {
		t.Error("Author is", post.Author)
	}
}
