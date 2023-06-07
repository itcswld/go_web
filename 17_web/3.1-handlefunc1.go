//handler
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// h.ServeHTTP(w, r)
		fmt.Printf("time:%s | handlerfunc:%s\n", time.Now().String(), runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name())
	}
}

func main() {
	h := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Gopher")
	}
	http.HandleFunc("/", log(h))
}
