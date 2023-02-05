//多處理器
package main

import (
	"fmt"
	"net/http"
)

type handle1 struct{}

func (h1 *handle1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handle1")
}

type handle2 struct{}

func (h2 *handle2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handle2")
}

func main() {
	h1 := handle1{}
	h2 := handle2{}

	svr := http.Server{
		Addr:    ":80",
		Handler: nil, //nil表示使用default多工器(DefaultServeMux)
	}
	//調用DefaultServeMux.Handle
	http.Handle("/h1", &h1) //localhost/h1
	http.Handle("/h2", &h2)
	svr.ListenAndServe()
}
