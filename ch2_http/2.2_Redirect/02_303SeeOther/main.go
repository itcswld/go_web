package main

import (
	"fmt"
	"html/template"
	"net/http"
)
var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("./*"))
}

func main(){
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

///303-see other
func b(w http.ResponseWriter, r *http.Request) {
	fmt.Println("b:", r.Method)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
// b: GET...http.Redirect(w, r, "/",)= localhost/a
// a: GET
// a: GET


func c(w http.ResponseWriter, r *http.Request) {
	fmt.Println("c:", r.Method)
	tpl.ExecuteTemplate(w,"index.html",nil)
}
// c: GET
// a: GET
//after submit... <form method="POST" action="/b"> = localhost/b
//b: POST...http.Redirect(w, r, "/",)= localhost/a
// a: GET
// a: GET
