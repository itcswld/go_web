//serverMux多工器
package main

import (
	"fmt"
	"log"
	"net/http"
)

func h1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Gopher")
}
func h2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Gopher2")
}

func main() {
	//way1
	// http.HandleFunc("/", hi)
	// if err := http.ListenAndServe(":80", nil); err != nil {
	// 	log.Fatal(err)
	// }

	//way2
	mux := http.NewServeMux()
	mux.HandleFunc("/h1", h1)
	mux.HandleFunc("/h2", h2)
	svr := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}
	if err := svr.ListenAndServe(); err != nil {
		log.Fatal(err)
		fmt.Println(err.Error())
	}
}
