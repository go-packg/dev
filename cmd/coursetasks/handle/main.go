package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func root(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<H1>You are on the top</H1>")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<H1>Bow-wow!</H1>")
}

func me(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(res, "<H1>My name is Slim Shadey!</H1>")
}

type myHandler struct {
	// ...
}

func (h myHandler) test(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<H1>My name is Valeri/ It's a test!</H1>")
}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.test(w, r)
}

func logchain(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before")
		io.WriteString(w, "<H1>Log::Before</H1>")
		h.ServeHTTP(w, r)
		io.WriteString(w, "<H1>Log::After</H1>")
		log.Println("After")
	})
}

func main() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(root))
	mux.Handle("/dog/", http.HandlerFunc(dog))
	//mux.Handle("/me/", http.HandlerFunc(me))
	//http.Handle("/me/", myHandler{})
	mux.Handle("/me/", myHandler{})
	mux.Handle("/test/", logchain(myHandler{}))
	http.ListenAndServe(":8080", mux)
}
