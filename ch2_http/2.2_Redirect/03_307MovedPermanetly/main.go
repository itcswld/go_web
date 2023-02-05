package main

import (
	"fmt"
	"html/template"
	"net/http"
)
var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	http.HandleFunc("/",a)
	http.HandleFunc("/b",b)
	http.HandleFunc("/c",c)
	http.ListenAndServe("",nil)
}

func a(w http.ResponseWriter, r *http.Request) {
	fmt.Println("a:", r.Method)
}
// a: GET
// a: GET

func b(w http.ResponseWriter, r *http.Request) {
	fmt.Println("b:", r.Method)
	// http.Redirect(w, r, "/", 307)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
//first time "/b"
// b: GET
// a: GET
// a: GET
//second time "/b"
// a: GET
// a: GET

func c(w http.ResponseWriter, r *http.Request){
	fmt.Println("c:", r.Method)
	tpl.ExecuteTemplate(w,"index.html",nil)
}
// c: GET
// a: GET
///...after submit = second time "/b"
// b: POST
// a: GET
// a: GET
