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

func main(){
	http.HandleFunc("/",foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request){
	//body
	b := make([]byte, r.ContentLength)
	r.Body.Read(b)
	body := string(b)
	err := tpl.ExecuteTemplate(w,"index.html", body)
	if err != nil {
		http.Error(w,err.Error(), 500)
		fmt.Print(err)
	}

}
