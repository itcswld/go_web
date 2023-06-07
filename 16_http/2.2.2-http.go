//http
package main

import (
	"fmt"
	"net/http"
	"reflect"
)

type handler bool

func (a handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {
	var a handler
	fmt.Println(reflect.TypeOf(a))
	http.ListenAndServe(":8080", a)
}

//curl http://localhost:8080
