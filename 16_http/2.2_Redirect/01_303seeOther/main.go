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
	http.HandleFunc("/", a)
	http.HandleFunc("/b",b)
	http.HandleFunc("/c",c)
	http.ListenAndServe("",nil)
}

func a(w http.ResponseWriter, r *http.Request){
	fmt.Println("a:",r.Method)
}
func b(w http.ResponseWriter, r *http.Request){
	fmt.Println("b:", r.Method)
	//process form data
	w.Header().Set("Location","/")
	w.WriteHeader(http.StatusSeeOther)
}

func c(w http.ResponseWriter, r *http.Request){
	fmt.Println("c:", r.Method)
	tpl.ExecuteTemplate(w, "index.html", nil)
}
