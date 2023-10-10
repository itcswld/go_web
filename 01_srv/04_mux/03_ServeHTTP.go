package main

import (
	"fmt"
	"log"
	"net/http"
)

type Acc struct {
	User string
}

//ServeHTTP 自訂處理器
func (a Acc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s , Welcome Home", a.User)
}

//處理函數
func Specpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The Go Programming Language Specification")
}

func main() {
	mux := http.NewServeMux()
	//註冊處理器
	mux.Handle("/", Acc{User: "Eve"})
	//註冊處理函數
	mux.HandleFunc("/spec", Specpage)

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
