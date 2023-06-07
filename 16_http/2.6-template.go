//template
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		//1. parse
		t, err := template.ParseFiles("./2.6-tpl.html")
		if err != nil {
			fmt.Println("tpl parse fail:", err.Error())
			return
		}
		//2.render
		name := "i love Golang"
		t.Execute(w, name)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
