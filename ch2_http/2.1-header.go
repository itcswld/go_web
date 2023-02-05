//response
package main

import (
	"fmt"
	"net/http"
)

func handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("token", "123456")
		w.Header().Set("Content-Type", "text/html; charset:utf-8")
		w.Header().Set("auther", "EveTong")
		w.Header().Set("date", "2022/2/15")
		fmt.Fprintln(w, "<h1>from .go</h1>")
	}
}

func main() {
	http.ListenAndServe(":8080", handler())
}

//回應狀態碼
//curl -i localhost:8080
