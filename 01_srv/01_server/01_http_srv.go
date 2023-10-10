package main

import "net/http"

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome Home."))
}

func main() {
	http.HandleFunc("/home", HomePage)
	http.ListenAndServe(":8080", nil)
}

//http://localhost:8080/home
