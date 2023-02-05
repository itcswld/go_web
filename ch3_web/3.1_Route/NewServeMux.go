package main

import (
	"io"
	"net/http"
)

func dog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "dog doggy")
	}
}
func  cat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "cat kitty")
	}
}
func main(){
	mux := http.NewServeMux()
	mux.Handle("/dog", dog())
	mux.Handle("/cat",cat())

	http.ListenAndServe(":8080", mux)
}
