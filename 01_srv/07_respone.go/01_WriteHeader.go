package main

import (
	"fmt"
	"net/http"
)

func UnAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
	fmt.Fprintln(w, http.StatusUnauthorized)
}

func main() {
	http.HandleFunc("/unAuth", UnAuth)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
