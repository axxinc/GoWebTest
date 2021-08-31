package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("temp.html")
	t.Execute(w, "Hello World")
}

func randNum(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("temp.html")

	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func rangeAction(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("temp.html")
	// daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, nil)
}

func withAction(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("temp.html")
	t.Execute(w, "hello")
}

func includeAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("t1.html", "t2.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, "Hello")
	fmt.Println("ok")
}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

// func funcMapAction(w http.ResponseWriter, r *http.Request) {
// 	funcMap := template.FuncMap{"fdate": formDate}
// 	t := template.New("tmpl.html").Funcs(funcMap)
// 	t, _ = template.ParseFiles("tmpl.html")
// 	t.Execute(w, time.Now())
// }
func funcMapAction(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("tmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("tmpl.html")
	t.Execute(w, time.Now())
}

func content(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("content.html")
	content := `I asked: <i>"what's up"</i>`
	t.Execute(w, content)
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}

func xxs(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("xxs.html")
	t.Execute(w, template.HTML(r.FormValue("comment")))
}

func layout(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html")
	t.ExecuteTemplate(w, "layout", "")
}

func layout2(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = t.ParseFiles("layout.html", "redhello.html")
	} else {
		// t, _ = t.ParseFiles("layout.html", "bulehello.html")
		// block
		t, _ = t.ParseFiles("layout.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/randNum", randNum)
	http.HandleFunc("/range", rangeAction)
	http.HandleFunc("/with", withAction)
	http.HandleFunc("/include", includeAction)

	http.HandleFunc("/funcMap", funcMapAction)

	http.HandleFunc("/content", content)

	http.HandleFunc("/form", form)
	http.HandleFunc("/xxs", xxs)

	http.HandleFunc("/layout", layout)
	http.HandleFunc("/layout2", layout2)
	// http.ListenAndServe(":8000", nil)
	server.ListenAndServe()
}
