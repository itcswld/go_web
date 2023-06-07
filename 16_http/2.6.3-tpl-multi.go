//template multipart
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	tpl = template.Must(template.ParseGlob("./*"))
)

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe("", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	f := r.FormValue("first")
	l := r.FormValue("last")
	s := r.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(w, "2.6.3-index.html", person{f, l, s})
	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Println(err)
	}
}
