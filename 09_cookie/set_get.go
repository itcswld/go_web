package main

import (
	"fmt"
	"net/http"
	"strconv"
)
var count = 0
func main(){
	http.HandleFunc("/",set)
	http.HandleFunc("/read",get)
	http.ListenAndServe("",nil)
}
func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "cookie"+strconv.Itoa(count),
		Value:   strconv.Itoa(count),
		Path: "/",
	})
	count +=1
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK BROWSER")
}

func get(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < count; i ++ {
			n := "cookie" + strconv.Itoa(i)
			c, err := r.Cookie(n)
		if err != nil {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}else{
			fmt.Fprintln(w, "cookie:", c)
		}
	}
}
//my-cookie=my-value
