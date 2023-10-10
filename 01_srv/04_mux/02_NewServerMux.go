package main

import (
	"fmt"
	"log"
	"net/http"
)

//NewServerMux 自訂Mulitplexer

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home")
}
func Specpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The Go Programming Language Specification")
}
func Aboutpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is an open-source programming language supported by Google")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Homepage) //http://localhost:8080/spec/ 會向上尋找 http://localhsot:8080/
	mux.HandleFunc("/spec", Specpage)
	mux.HandleFunc("/spec/intro", Aboutpage)

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
