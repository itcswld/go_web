//enctype = "application/x-www-form-urlencoded"， 用於獲取表單值
package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, "鍵值對和url鍵值對 : ", r.Form)
	fmt.Fprintln(w, "鍵值對：", r.PostForm)
}

func main() {
	http.HandleFunc("/", process)
	http.ListenAndServe(":8080", nil)
}
