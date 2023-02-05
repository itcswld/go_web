//handler = handlerFunc()
package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello from a HandleFunc #1!\n")
}
func cat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}
}
func main() {
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", cat())

	http.ListenAndServe(":8080", nil)
}
