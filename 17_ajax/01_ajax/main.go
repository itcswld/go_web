package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/text", foo)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	s := `this text  is from Go`
	fmt.Fprintln(w, s)
}
