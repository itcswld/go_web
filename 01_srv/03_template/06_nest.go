package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./06_Nhome.html", "./06_Nbody.html", "./06_Nfooter.html")
	if err != nil {
		fmt.Println("tpl err: ", err)
		return
	}
	user := struct {
		App, Author, User string
		Types             []string
	}{"Go", "Eve", "Jay", []string{"Boolean", "Numeric", "String", "Array", "Slice", "Struct"}}
	tpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/home", HomePage)
	http.ListenAndServe(":8080", nil)
}
