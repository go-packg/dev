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

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`hello world`))
	io.WriteString(w, "<H1>>My name is Valeri!</H1>")
}

func logchain(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before")
		h.ServeHTTP(w, r) // call original
		log.Println("After")
	})
}

func test(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "<H1>My name is Valeri/ It's a test!</H1>")
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
	mux.Handle("/test/", logchain(test))
	http.ListenAndServe(":8080", mux)
}
