package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is %s", r.Proto) //this is HTTP/2.0
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServeTLS(":80", "cert.pem", "key.pem", nil))
}

//https://localhost:80/
