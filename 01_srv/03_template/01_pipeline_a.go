package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./01_pipeline_a.html")
	if err != nil {
		fmt.Println(err)
	}
	v := "I love coding."
	tpl.Execute(w, v)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
