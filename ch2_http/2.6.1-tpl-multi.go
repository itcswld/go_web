//template multipart
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Name, Phone string
}

func tpl(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./2.6.1-tpl-multi.html", "./2.6.1-ul.html")
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	user := Info{"Eve", "09123432"}
	tpl.Execute(w, user)
	fmt.Println(tpl)
}
func main() {
	http.HandleFunc("/", tpl)
	http.ListenAndServe(":80", nil)
}
