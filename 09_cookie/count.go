//creating a counter to track how many times a user has been to a particular website
package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/",set)
	http.Handle("/a", http.NotFoundHandler())
	http.ListenAndServe("",nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("cookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name: "cookie",
			Value: "0",
			Path: "/",
		}
	}
	//atoi: string to int
	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Error(w,http.StatusText(400), http.StatusBadRequest)
		return
	}
	count++
	//Itoa: int to string
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)
}
