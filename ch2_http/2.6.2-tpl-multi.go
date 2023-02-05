//template multipart
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"
)

func handler() http.HandlerFunc {
	tpl := template.Must(template.ParseFiles("2.6.2-tpl-multi.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		fmt.Println(reflect.TypeOf(r.Form)) //url.Values
		tpl.Execute(w, r.Form)
	}
}

func main() {
	http.ListenAndServe(":8080", handler())
}
