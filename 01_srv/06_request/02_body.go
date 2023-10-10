package main

import (
	"fmt"
	"net/http"
)

func getBody(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	fmt.Println(len) //2
	r.Body.Read(body)
	w.Write(body)
}

func main() {
	http.HandleFunc("/getbody", getBody)
	http.ListenAndServe(":8080", nil)
}

//curl http://localhost:8080/getbody -X POST -d "go"
