package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Info struct {
	ID   string
	Name string
	Life int
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := Info{
		ID:   "E08888",
		Name: "PC",
		Life: 3,
	}
	tpl, err := template.ParseFiles("./01_pipeline_b.html")
	if err != nil {
		fmt.Println(err)
	}
	tpl.Execute(w, data)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
