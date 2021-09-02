package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

// type PostTestSuite struct{}

type PostTestSuite struct {
	mux   *http.ServeMux
	post  *Post
	write *httptest.ResponseRecorder
}

func init() {
	Suite(&PostTestSuite{})
}

func (s *PostTestSuite) SetUpTest(c *C) {
	var err error
	Db, err := sql.Open("postgres", "user=postgres password=123456 port=5432 host=127.0.0.1 sslmode=disable dbname=gwp")
	if err != nil {
		panic(err)
	}

	s.post = &Post{Db: Db}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/post/", handleRequest(s.post))
	s.write = httptest.NewRecorder()
}

func (s *PostTestSuite) TestHandleGet(c *C) {

	// mux := http.NewServeMux()
	// mux.HandleFunc("/post/", handleRequest(&Post{Db: Db}))

	// write := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/post/1", nil)

	s.mux.ServeHTTP(s.write, request)

	c.Check(s.write.Code, Equals, 200)

	var post Post
	json.Unmarshal(s.write.Body.Bytes(), &post)

	c.Check(post.Id, Equals, 1)
}

func (s *PostTestSuite) TestHandlePost(c *C) {
	json := strings.NewReader(`{"content":"11111","author":"zhangsan"}`)

	request, _ := http.NewRequest("PUT", "/post/1", json)
	s.mux.ServeHTTP(s.write, request)

	c.Check(s.write.Code, Equals, 200)
	c.Check(s.post.Id, Equals, 1)
	c.Check(s.post.Author, Equals, "zhangsan")
}

func Test(t *testing.T) {
	TestingT(t)
}

func (s *PostTestSuite) TearDownTest(c *C) {
	c.Log("Finished test - ", c.TestName())
}

func (s *PostTestSuite) SetUpSuite(c *C) {
	c.Log("Starting Test Suite ")
}

func (s *PostTestSuite) TearDownSuite(c *C) {
	c.Log("Finished Test Suite ")
}
