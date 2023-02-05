package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main(){
	http.HandleFunc("/", set)
	http.HandleFunc("/read", get)
	http.ListenAndServe("",nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	for i :=0; i < 2; i++{
		no :=  strconv.Itoa(i+1)
		http.SetCookie(w, &http.Cookie{
			Name: "cookie" + no,
			Value: "value1" + no,
		})
	}
	fmt.Fprintln(w, "COOKIES WRITTEN - CHECK BROWSER")
}

func get(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 2; i ++ {
		c := "cookie" + strconv.Itoa(i)
		cookie, err := r.Cookie(c)
		if err != nil {
			http.Error(w, http.StatusText(400),http.StatusBadRequest)
			return
		}else{
			fmt.Fprintln(w,"",cookie)
		}
	}
}
