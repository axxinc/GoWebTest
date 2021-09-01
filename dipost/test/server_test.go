package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)

	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Respose code is %v", writer.Code)
	}

	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&FakePost{}))

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"Update post","author":"sansan"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)

	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Error("Response code is ", writer.Code)
	}

	// post, _ := retrieve(1)
	// if post.Content != "Update post" {
	// 	t.Error("Post content is ", post.Content)
	// }
}
