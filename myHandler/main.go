package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

type myHandler struct{}

func (myhandler myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello myHandler")
}

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

type worldHandler struct{}

func (world *worldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello1")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world1")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called -" + name)
		h(w, r)
	}
}

func header(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func process(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// fmt.Fprintln(w, r.Form) // urlencoded: map[firstname:[4 2] hello:[3 1]]    form-data: map[firstname:[2] hello:[1]]
	// fmt.Fprintln(w, r.PostForm) // urlencoded: map[firstname:[4] hello:[3]]

	// r.ParseMultipartForm(1024)
	// fmt.Fprintln(w, r.MultipartForm) // form-data: &{map[firstname:[4] hello:[3]] map[]}

	// fmt.Fprintln(w, r.FormValue("hello"))     // 3
	// fmt.Fprintln(w, r.PostFormValue("hello")) // 3

	// fmt.Fprintln(w, r.Form) // map[firstname:[4 2] hello:[3 1]]

	// r.ParseMultipartForm(1024)
	// fileHeader := r.MultipartForm.File["upload"]
	// file, err := fileHeader[0].Open()

	file, _, err := r.FormFile("upload")

	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func write(w http.ResponseWriter, r *http.Request) {
	str := `<html> 
	<head>
	<title>Go Web Programming</title>
	</head>
	<body>aaaaa</body>`

	w.Write([]byte(str))
}

func writeHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintf(w, "no such service,try next door ")
}

func wHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("location", "http://www.baidu.com")
	w.WriteHeader(302)
}

type Post struct {
	User    string
	Threads []string
}

func JsonWrite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	post := &Post{
		User:    "zhangsan",
		Threads: []string{"111", "222", "333"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Maning Publication Go",
		HttpOnly: true,
	}

	// w.Header().Set("Set-Cookie", c1.String())
	// w.Header().Add("Set-Cookie", c2.String())

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header["Cookie"]
	fmt.Fprintln(w, cookie)

	firstname, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Println("cannot get first Cookie")
	}
	fmt.Fprintln(w, firstname)

	cookies := r.Cookies()
	fmt.Fprintln(w, cookies)
}

func setMessage(w http.ResponseWriter, r *http.Request) {
	message := []byte("Hello World")
	flash := http.Cookie{
		Name:  "message",
		Value: base64.URLEncoding.EncodeToString(message),
	}
	http.SetCookie(w, &flash)
}
func getMessage(w http.ResponseWriter, r *http.Request) {
	flash, err := r.Cookie("message")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "no message found")
		}
	} else {
		rc := http.Cookie{
			Name:    "message",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)

		message, _ := base64.URLEncoding.DecodeString(flash.Value)
		fmt.Fprintln(w, string(message))

	}
}

func main() {
	// handler := myHandler{}
	h := helloHandler{}
	w := worldHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8000",
		// Handler: &handler,
	}
	http.Handle("/hello", &h)
	http.Handle("/world", &w)

	http.HandleFunc("/hello1", hello)
	http.HandleFunc("/world1", world)

	http.HandleFunc("/hello2", log(hello))

	http.HandleFunc("/header", header)
	http.HandleFunc("/body", body)

	http.HandleFunc("/process", process)

	http.HandleFunc("/write", write)
	http.HandleFunc("/writeHeader", writeHeader)
	http.HandleFunc("/wHeader", wHeader)
	http.HandleFunc("/JsonWrite", JsonWrite)

	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)

	http.HandleFunc("/setMessage", setMessage)
	http.HandleFunc("/getMessage", getMessage)

	// server.ListenAndServeTLS("cert.pem", "key.pem")
	server.ListenAndServe()
}
